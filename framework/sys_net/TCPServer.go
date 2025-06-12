package sys_net

import (
	"encoding/binary"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"net"
	"ricebean/framework/log4"
	"ricebean/framework/sys_map"
	"runtime/debug"
)

type Server struct{}

type IServer interface {
	ServeForever()
	KillConnect(netId string, prefix string)
	SendMessage(netId string, msgId int32, msgData []byte) error
	CloseServer()
	GetConnectNum() int32
}

type ICTXSession interface {
	CheckCheat() bool
	UpdateConnectTime()
}

type ICTXSessionManager interface {
	SetServer(server IServer)
	AddConnect(netId string)
	GetConnect(netId string) *ICTXSession
	RemoveConnect(netId string)
}

type IServerEventHandler interface {
	Init(t IServer)
	FailHandler(t IServer, prefix string) //监听失败
	ConnAddHandler(netId string)          //添加连接
	ConnCloseHandler(netId string)        //断开连接
}

type ServerEventHandler struct {
	Init             func(t IServer)
	HandlerData      func(netId string, msgId int32, msgData []byte)
	FailHandler      func(t IServer, prefix string)
	ConnAddHandler   func(netId string)
	ConnCloseHandler func(netId string)
}

type TCPServer struct {
	_connMap          *sys_map.ConcurrentMap[string, net.Conn]
	name              string
	ipaddress         string
	workerPool        *goroutine.Pool
	_listener         *net.TCPListener
	_HandlerData      func(netId string, msgId int32, msgData []byte)
	_FailHandler      func(t IServer, prefix string)
	_ConnAddHandler   func(netId string)
	_ConnCloseHandler func(netId string)
	dataChan          chan *DataPackage
}

type DataPackage struct {
	netId   string
	msgId   int32
	msgData []byte
}

func NewTCPServer(ipaddress string,
	HandlerData func(netId string, msgId int32, msgData []byte),
	FailHandler func(t IServer, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string),
) *TCPServer {
	t := &TCPServer{}
	t.ipaddress = ipaddress
	t._connMap = sys_map.NewConcurrentMap[string, net.Conn]()
	t.workerPool = goroutine.Default()
	t._HandlerData = HandlerData
	t._FailHandler = FailHandler
	t._ConnAddHandler = ConnAddHandler
	t._ConnCloseHandler = ConnCloseHandler
	t.dataChan = make(chan *DataPackage, 100000) // 创建带缓冲的channel
	return t
}

func (t *TCPServer) ServeForever() {
	t._connMap = sys_map.NewConcurrentMap[string, net.Conn]()
	addr, err := net.ResolveTCPAddr("tcp", t.ipaddress)
	if err != nil {
		t._FailHandler(t, "IpAddress is Error")
		return
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		t._FailHandler(t, "context fail"+err.Error())
		return
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

func (t *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	t._OnOpened(conn)
	byteBuffer := make([]byte, 0)
	buf := make([]byte, 1024)
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
				t.dataChan <- &DataPackage{netId: netId, msgId: msgId, msgData: content}
			}
		}

	}
}

func (t *TCPServer) processData() {
	for dataPkg := range t.dataChan {
		defer func() {
			if r := recover(); r != nil {
				er := debug.Stack()
				log4.Error("处理数据异常 msgId=%d:%s %s", dataPkg.msgId, r, er)
			}
		}()

		t._HandlerData(dataPkg.netId, dataPkg.msgId, dataPkg.msgData)
	}
}

func (t *TCPServer) _OnOpened(conn net.Conn) {
	netId := conn.RemoteAddr().String()
	t._connMap.Set(netId, conn)
	log4.Info("accept -ctx->netId  %s-", netId)
	t._ConnAddHandler(netId)
}

func (t *TCPServer) _OnClosed(conn net.Conn, prefix string) {
	netId := conn.RemoteAddr().String()
	t._connMap.Delete(netId)
	log4.Info("closed -ctx->netId  %s-", netId)
	t._ConnCloseHandler(netId)
}

func (t *TCPServer) GetConn(netId string) net.Conn {
	return t._connMap.Get(netId)
}

func (t *TCPServer) KillConnect(netId string, prefix string) {
	conn := t._connMap.Get(netId)
	if conn == nil {
		return
	}
	conn.Close()
	conn = nil
	t._connMap.Delete(netId)
}
func (t *TCPServer) KillAllConnect(prefix string) {
	t._connMap.Range(func(netId string, conn net.Conn) bool {
		t._connMap.Delete(netId)
		if conn != nil {
			conn.Close()
			conn = nil
		}
		return true
	})
}
func (t *TCPServer) SendMessage(netId string, msgId int32, msgData []byte) error {
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

func (t *TCPServer) GetConnectNum() int32 {
	return int32(t._connMap.Len())
}

func (t *TCPServer) CloseServer() {
	if t._listener != nil {
		t._listener.Close()
		t._listener = nil
	}
}
