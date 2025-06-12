package sys_net

import (
	"encoding/binary"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"github.com/xtaci/kcp-go"
	"net"
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/framework/sys_map"
)

/*
	KCP
added by yh @ 2023/5/29 16:22
注意:
*/

type KCPServer struct {
	_connMap          *sys_map.NormalMap[string, net.Conn]
	name              string
	ipaddress         string
	workerPool        *goroutine.Pool
	_listener         net.Listener
	_HandlerData      func(netId string, msgId int32, msgData []byte)
	_FailHandler      func(t IServer, prefix string)
	_ConnAddHandler   func(netId string)
	_ConnCloseHandler func(netId string)
	dataChan          chan *DataPackage
}

/*
type DataPackage struct {
	netId   string
	msgId   int32
	msgData []byte
}*/

func NewKCPServer(ipaddress string,
	Init func(t IServer),
	HandlerData func(netId string, msgId int32, msgData []byte),
	FailHandler func(t IServer, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string),
) *KCPServer {
	t := &KCPServer{}
	t.ipaddress = ipaddress
	t._connMap = sys_map.NewNormalMap[string, net.Conn]()
	t.workerPool = goroutine.Default()
	t._HandlerData = HandlerData
	t._FailHandler = FailHandler
	t._ConnAddHandler = ConnAddHandler
	t._ConnCloseHandler = ConnCloseHandler
	t.dataChan = make(chan *DataPackage, 1000000) // 创建带缓冲的channel
	return t
}

func (t *KCPServer) ServeForever() {
	t._connMap = sys_map.NewNormalMap[string, net.Conn]()
	// 创建 KCP 侦听器 ":9999"
	listener, err := kcp.Listen(t.ipaddress)
	if err != nil {
		t._FailHandler(t, "context fail")
	}
	defer listener.Close()

	t._listener = listener

	// 启动数据处理 goroutine
	go t.processData()

	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			conn.Close()
			conn = nil
			continue
		}
		//go t.handleConnection(conn)
		_ = t.workerPool.Submit(func() {
			t.handleConnection(conn)
		})

	}
}

func (t *KCPServer) CloseServer() {
	if t._listener != nil {
		t._listener.Close()
		t._listener = nil
	}
}

func (t *KCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	t._OnOpened(conn)
	byteBuffer := make([]byte, 0)
	buf := make([]byte, 512)
	bodyLen := 0

	for {
		n, err := conn.Read(buf)
		if err != nil {
			t._OnClosed(conn, "Error reading from connection")
			return
		}
		if n == 0 {
			continue
		}
		byteBuffer = append(byteBuffer, buf[:n]...)
		if len(byteBuffer) < 4 {
			continue
		}
		bodyLen = int(binary.LittleEndian.Uint32(byteBuffer[:4]))
		if bodyLen > p_buffer_len {
			conn.Close()
			byteBuffer = make([]byte, 0)
			t._OnClosed(conn, "bodyLen > p_buffer_len")
			return
		}
		if len(byteBuffer) >= bodyLen+4 {
			data := byteBuffer[0 : bodyLen+4]
			byteBuffer = byteBuffer[bodyLen+4:]

			msgId := int32(binary.LittleEndian.Uint32(data[4:8]))
			content := data[8:]
			netId := conn.RemoteAddr().String()

			//CSP（Communicating Sequential Processes） 通信顺序进程 一种并发模式
			t.dataChan <- &DataPackage{netId: netId, msgId: msgId, msgData: content}
		}
	}
}

func (t *KCPServer) processData() {
	for dataPkg := range t.dataChan {

		/*
			defer func() {
				if r := recover(); r != nil {
					er := debug.Stack()
					log4.Error("处理数据异常 msgId=%d:%s %s", dataPkg.msgId, r, er)
				}
			}()
		*/
		t._HandlerData(dataPkg.netId, dataPkg.msgId, dataPkg.msgData)
	}
}

func (t *KCPServer) _OnOpened(conn net.Conn) {
	netId := conn.RemoteAddr().String()
	t._connMap.Set(netId, conn)
	log4.Info("accept -ctx->netId  %s-", netId)
	t._ConnAddHandler(netId)
}

func (t *KCPServer) _OnClosed(conn net.Conn, prefix string) {
	netId := conn.RemoteAddr().String()
	t._connMap.Delete(netId)
	log4.Info("closed -ctx->netId  %s-", netId)
	t._ConnCloseHandler(netId)
}

func (t *KCPServer) GetConn(netId string) net.Conn {
	return t._connMap.Get(netId)
}

func (t *KCPServer) KillConnect(netId string, prefix string) {
	conn := t._connMap.Get(netId)
	if conn == nil {
		return
	}
	conn.Close()
	conn = nil
	t._connMap.Delete(netId)
}

func (t *KCPServer) SendMessage(netId string, msgId int32, msgData []byte) error {
	conn := t._connMap.Get(netId)
	if conn == nil {
		return nil
	}
	msgIdBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgIdBytes, uint32(msgId))
	content := append(msgIdBytes, msgData...)
	var l = int32(len(content))
	lByte := make([]byte, 4)
	binary.LittleEndian.PutUint32(lByte, uint32(l))
	msgPkg := append(lByte, content...)

	_, err := conn.Write(msgPkg)
	return err
}
func (t *KCPServer) GetConnectNum() int32 {
	//TODO implement me
	return 0
}
