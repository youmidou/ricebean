package sys_net

import (
	"encoding/binary"
	"net"
	"phoenix-tudou/framework/log4"
	"runtime/debug"
)

/*
	TCP 客户端

added by yh @ 2024/04/25 12:57
注意:
*/
type TCPClient struct {
	ipaddress         string
	_HandlerData      func(netId string, msgId int32, msgData []byte)
	_FailHandler      func(c *TCPClient, prefix string) //连接失败
	_ConnAddHandler   func(netId string)
	_ConnCloseHandler func(netId string)

	conn net.Conn
}

func NewTCPClient(
	IpAddress string,
	HandlerData func(netId string, reqId int32, bytes []byte),
	FailHandler func(c *TCPClient, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string)) *TCPClient {
	t := &TCPClient{}
	t.ipaddress = IpAddress
	t._HandlerData = HandlerData
	t._FailHandler = FailHandler
	t._ConnAddHandler = ConnAddHandler
	t._ConnCloseHandler = ConnCloseHandler
	return t
}
func (t *TCPClient) ConnectForever() {
	// 连接到服务器
	conn, err := net.Dial("tcp", t.ipaddress)
	if err != nil {
		//log4.Info("tcp client dail tcp ipaddress: %s %s", t.ipaddress, err.Error())
		t._FailHandler(t, "context fail")
		return
	}
	defer conn.Close()
	t._OnOpened(conn)
	byteBuffer := make([]byte, 0)
	bodyLen := 0
	// 读取客户端发送的数据
	buf := make([]byte, 512)
	// 循环读取数据
	for {
		n, err2 := conn.Read(buf)
		if err2 != nil {
			defer func() {
				if r := recover(); r != nil {
					er := debug.Stack()
					log4.Error("客户端主动断开连接->发生错误抛出 :%s %s", r, er)
				}
			}()
			t._OnClosed(conn, "Error reading from connection")
			t._FailHandler(t, "Error reading from connection")
			return
		}
		if n <= 0 {
			continue
		}
		// 将接收到的数据丢进缓存池子里去
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
			//if len(byteBuffer) >= bodyLen {
			data := byteBuffer[4 : bodyLen+4]
			byteBuffer = byteBuffer[bodyLen+4:]
			//-------------------------------------
			msgId := int32(binary.LittleEndian.Uint32(data[0:4]))
			content := data[4:]
			netId := conn.RemoteAddr().String()

			defer func() {
				if r := recover(); r != nil {
					er := debug.Stack()
					log4.Error("处理数据异常 msgId=%d:%s %s", msgId, r, er)
				}
			}()
			// 处理数据
			t._HandlerData(netId, msgId, content)
			//CSP（Communicating Sequential Processes） 通信顺序进程 一种并发模式
			//t.dataChan <- &DataPackage{netId: netId, msgId: msgId, msgData: content}
			//}
		}
	}
}

/*
// 握手
func (t *TCPClient) DoHandshake() error {

	PacketEncoder := codec.NewPomeloPacketEncoder()
	msgPkg2, e := PacketEncoder.Encode(packet.Data, msgPkg)
	if e != nil {
		return e
	}
	_, err := t.conn.Write(msgPkg2)
	return err
}
*/

func (t *TCPClient) SendMessage(msgId int32, msgData []byte) error {
	// 发送数据
	if t.conn == nil {
		return net.ErrClosed
	}
	//将 msgId 转换成[]byte
	msgIdBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgIdBytes, uint32(msgId))

	// 将 msgIdBytes 和 content 组合成一个新的 []byte 切片 result
	content := append(msgIdBytes, msgData...)
	l := int32(len(content))
	l_btye := make([]byte, 4)
	binary.LittleEndian.PutUint32(l_btye, uint32(l))
	//--封包---
	msgPkg := append(l_btye, content...)
	_, err := t.conn.Write(msgPkg)
	return err
}

func (t *TCPClient) _OnOpened(conn net.Conn) {
	netId := conn.RemoteAddr().String()
	t.conn = conn
	log4.Info("accept -ctx->netId  %s-", netId)
	if t._ConnAddHandler != nil {
		t._ConnAddHandler(netId)
	}

}
func (t *TCPClient) _OnClosed(conn net.Conn, prefix string) {
	netId := conn.RemoteAddr().String()
	t.conn = nil
	log4.Info("closed -ctx->netId  %s-", netId)
	if t._ConnCloseHandler != nil {
		t._ConnCloseHandler(netId)
	}
}

func (t *TCPClient) UpdataTime() {

}

func (t *TCPClient) GetNetId() string {
	if t.conn != nil {
		return t.conn.RemoteAddr().String()
	}
	return "-1"
}
