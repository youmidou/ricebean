package sys_net

import (
	"encoding/binary"
	"github.com/xtaci/kcp-go"
	"net"
	"ricebean/framework/log4"

	"runtime/debug"
)

/*
	KCP 客户端

added by yh @ 2024/04/25 12:57
注意:
*/
type KCPClient struct {
	ipaddress         string
	_HandlerData      func(c *KCPClient, netId string, msgId int32, msgData []byte)
	_FailHandler      func(c *KCPClient, prefix string) //连接失败
	_ConnAddHandler   func(netId string)
	_ConnCloseHandler func(netId string)

	conn net.Conn
}

func NewKCPClient(IpAddress string, HandlerData func(c *KCPClient, netId string, reqId int32, bytes []byte),
	FailHandler func(c *KCPClient, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string),
) *KCPClient {
	t := &KCPClient{}
	t.ipaddress = IpAddress
	t._HandlerData = HandlerData
	t._FailHandler = FailHandler
	t._ConnAddHandler = ConnAddHandler
	t._ConnCloseHandler = ConnCloseHandler
	return t
}
func (t *KCPClient) ConnectForever() {
	// 连接到服务器
	conn, err := kcp.Dial("127.0.0.1:9999")
	if err != nil {
		//log4.Info("tcp client dail tcp ipaddress: %s %s", t.ipaddress, err.Error())
		t._FailHandler(t, "context fail")
		return
	}
	defer conn.Close()

	t._OnOpened(conn)
	byte_buffer := make([]byte, 0)
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
		byte_buffer = append(byte_buffer, buf[:n]...)
		if len(byte_buffer) < 4 {
			continue
		}
		bodyLen = int(binary.LittleEndian.Uint32(byte_buffer[:4]))
		if bodyLen > p_buffer_len {
			//错误消息断开连接
			conn.Close()
			t._OnClosed(conn, "bodyLen > p_buffer_len")
			t._FailHandler(t, "bodyLen > p_buffer_len")
			return
		}
		if len(byte_buffer) >= bodyLen+4 {
			data := byte_buffer[0 : bodyLen+4]
			byte_buffer = byte_buffer[bodyLen+4:] //截取剩余的数据

			// 解析消息号字段
			msgId := int32(binary.LittleEndian.Uint32(data[4:8]))
			// 解析内容
			content := data[8:]
			netId := conn.RemoteAddr().String()

			defer func() {
				if r := recover(); r != nil {
					er := debug.Stack()
					log4.Error("处理数据异常 msgId=%d:%s %s", msgId, r, er)
				}
			}()
			// 处理数据
			t._HandlerData(t, netId, msgId, content)
		}

	}
}

func (t *KCPClient) SendMessage(msgId int32, msgData []byte) error {
	//将 msgId 转换成[]byte
	msgIdBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgIdBytes, uint32(msgId))

	// 将 msgIdBytes 和 content 组合成一个新的 []byte 切片 result
	content := append(msgIdBytes, msgData...)
	var l int32 = int32(len(content))
	l_btye := make([]byte, 4)
	binary.LittleEndian.PutUint32(l_btye, uint32(l))
	//--封包---
	msgPkg := append(l_btye, content...)

	// 发送数据
	if t.conn == nil {
		return nil
	}
	_, err := t.conn.Write(msgPkg)
	return err
}

func (t *KCPClient) _OnOpened(conn net.Conn) {
	netId := conn.RemoteAddr().String()
	t.conn = conn
	log4.Info("accept -ctx->netId  %s-", netId)
	if t._ConnAddHandler != nil {
		t._ConnAddHandler(netId)
	}

}
func (t *KCPClient) _OnClosed(conn net.Conn, prefix string) {
	netId := conn.RemoteAddr().String()
	t.conn = nil
	log4.Info("closed -ctx->netId  %s-", netId)
	if t._ConnCloseHandler != nil {
		t._ConnCloseHandler(netId)
	}
}

func (t *KCPClient) UpdataTime() {

}

func (t *KCPClient) GetNetId() string {
	if t.conn != nil {
		return t.conn.RemoteAddr().String()
	}
	return "-1"
}
