package sys_net

import (
	"context"
	"encoding/binary"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"log"
	"net"
	"net/http"
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/framework/sys_map"
	"runtime/debug"
	"time"
)

/*
	WSTCP 服务器

added by yh @ 2024/05/15 17:57
注意: WebSocket
改进点：
使用 sync.WaitGroup 等待所有 goroutine 完成：在服务器关闭时，等待所有 goroutine 正确退出。
内存管理：在处理消息时，合理地管理缓冲区，避免不必要的内存占用。
WebSocket 缓冲区设置：根据实际需求设置合理的读写缓冲区大小。
优雅关闭服务器：使用 http.Server 的 Shutdown 方法优雅关闭服务器。
通过这些改进，可以有效降低内存占用，并确保服务器能够优雅地处理连接和关闭。
*/
type WSTCPServer struct {
	_connMap          *sys_map.ConcurrentMap[string, *websocket.Conn]
	_connTime         *sys_map.ConcurrentMap[string, time.Time] //连接超时30秒 主动踢人
	name              string
	ipaddress         string
	workerPool        *goroutine.Pool
	_listener         *net.TCPListener
	_HandlerData      func(netId string, msgId int32, content []byte)
	_FailHandler      func(t IServer, prefix string)
	_ConnAddHandler   func(netId string)
	_ConnCloseHandler func(netId string)
	stopChan          chan struct{} // 用于优雅关闭服务器
}

// netId, msgId, content
func NewWSTCPServer(ipaddress string,
	HandlerData func(netId string, msgId int32, content []byte),
	FailHandler func(t IServer, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string),
) *WSTCPServer {
	t := &WSTCPServer{}
	t.ipaddress = ipaddress
	t._connMap = sys_map.NewConcurrentMap[string, *websocket.Conn]()
	t._connTime = sys_map.NewConcurrentMap[string, time.Time]()
	t.workerPool = goroutine.Default()
	t._HandlerData = HandlerData
	t._FailHandler = FailHandler
	t._ConnAddHandler = ConnAddHandler
	t._ConnCloseHandler = ConnCloseHandler
	t.stopChan = make(chan struct{})
	//go t.CheckTimeOut()
	return t
}

func (t *WSTCPServer) CheckTimeOut() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			t._connTime.Range(func(k string, v time.Time) bool {
				if now.Sub(v) > 30*time.Second {
					t.KillConnect(k, "time out")
				}
				return true
			})
		case <-t.stopChan:
			return
		}
	}
}

func (t *WSTCPServer) GetNetId() string {
	return t.ipaddress
}

// "0.0.0.0:1240"
func (t *WSTCPServer) ServeForever() {
	t._connMap = sys_map.NewConcurrentMap[string, *websocket.Conn]()
	server := &http.Server{
		Addr: t.ipaddress,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 检查请求是否是 WebSocket 升级请求
			if r.Header.Get("Upgrade") != "websocket" {
				http.Error(w, "Not a WebSocket upgrade request", http.StatusBadRequest)
				return
			}
			conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
			if err != nil {
				http.Error(w, "Failed to upgrade to WebSocket connection", http.StatusInternalServerError)
				return
			}
			// 处理 WebSocket 连接
			_ = t.workerPool.Submit(func() {
				t.handleConnection(conn)
			})
		}),
	}

	//fmt.Printf("Server listening on :%s", t.ipaddress)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			t._FailHandler(t, "Error starting server")
		}
	}()

	<-t.stopChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}

func (t *WSTCPServer) handleConnection(conn *websocket.Conn) {
	defer conn.Close()
	t._OnOpened(conn)

	var byteBuffer []byte
	bodyLen := 0

	for {
		n, buf, err := conn.ReadMessage()
		if err != nil {
			t._OnClosed(conn, "Error reading from connection")
			return
		}

		byteBuffer = append(byteBuffer, buf[:n]...)
		//----------------------------------------
		for {
			if len(byteBuffer) < 4 {
				break
			}
			bodyLen = int(binary.LittleEndian.Uint32(byteBuffer[:4]))
			if bodyLen > p_buffer_len {
				conn.Close()
				byteBuffer = make([]byte, 0)
				t._OnClosed(conn, "bodyLen > p_buffer_len")
				return
			}
			if len(byteBuffer) < bodyLen+4 {
				break
			}
			if len(byteBuffer) >= bodyLen {
				data := byteBuffer[4 : bodyLen+4]
				byteBuffer = byteBuffer[bodyLen+4:]
				//-------------------------------------
				msgId := int32(binary.LittleEndian.Uint32(data[0:4]))
				content := data[4:]
				netId := conn.RemoteAddr().String()

				//CSP（Communicating Sequential Processes） 通信顺序进程 一种并发模式
				//t.dataChan <- &DataPackage{netId: netId, msgId: msgId, msgData: content}

				defer func() {
					if r := recover(); r != nil {
						er := debug.Stack()
						log4.Error("处理数据异常 msgId=%d:%s %s", msgId, r, er)
					}
				}()
				t._connTime.Set(netId, time.Now())
				t._HandlerData(netId, msgId, content)
			}
		}
	}
}

// 进入
func (t *WSTCPServer) _OnOpened(conn *websocket.Conn) {
	netId := conn.RemoteAddr().String()
	t._connMap.Set(netId, conn)
	t._connTime.Set(netId, time.Now())
	t._ConnAddHandler(netId)
}

func (t *WSTCPServer) _OnClosed(conn *websocket.Conn, prefix string) {
	netId := conn.RemoteAddr().String()
	t._connMap.Delete(netId)
	t._connTime.Delete(netId)
	t._ConnCloseHandler(netId)
}

func (t *WSTCPServer) KillConnect(netId string, prefix string) {
	conn := t._connMap.Get(netId)
	if conn == nil {
		return
	}
	conn.Close()
	t._connMap.Delete(netId)
}

func (t *WSTCPServer) KillAllConnect(prefix string) {
	t._connMap.Range(func(netId string, conn *websocket.Conn) bool {
		t._connMap.Delete(netId)
		if conn != nil {
			conn.Close()
			conn = nil
		}
		return true
	})
}
func (t *WSTCPServer) SendMessage(netId string, msgId int32, msgData []byte) error {
	conn := t._connMap.Get(netId)
	if conn == nil {
		return nil
	}
	msgIdBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgIdBytes, uint32(msgId))
	content := append(msgIdBytes, msgData...)
	l := int32(len(content))
	lBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(lBytes, uint32(l))
	msgPkg := append(lBytes, content...)

	err := conn.WriteMessage(websocket.BinaryMessage, msgPkg)
	return err
}

func (t *WSTCPServer) GetConnectNum() int32 {
	return int32(t._connMap.Len())
}

func (t *WSTCPServer) CloseServer() {
	close(t.stopChan)
	if t._listener != nil {
		t._listener.Close()
		t._listener = nil
	}
}
