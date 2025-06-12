package sys_net

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"ricebean/framework/log4"
	"runtime/debug"
)

/*
	TCP 客户端

added by yh @ 2024/05/15 18:28
注意:
*/
type WSTCPClient struct {
	ipaddress         string
	_HandlerData      func(netId string, msgId int32, msgData []byte)
	_FailHandler      func(c *WSTCPClient, prefix string) //连接失败
	_ConnAddHandler   func(netId string)
	_ConnCloseHandler func(netId string)

	conn *websocket.Conn
}

func NewWSTCPClient(IpAddress string,
	HandlerData func(netId string, reqId int32, bytes []byte),
	FailHandler func(c *WSTCPClient, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string),
) *WSTCPClient {
	t := &WSTCPClient{}
	t.ipaddress = IpAddress
	t._HandlerData = HandlerData
	t._FailHandler = FailHandler
	t._ConnAddHandler = ConnAddHandler
	t._ConnCloseHandler = ConnCloseHandler
	return t
}
func (t *WSTCPClient) ConnectForever() {
	// 连接到服务器

	// 设置 WebSocket 服务器地址
	//addr := "ws://0.0.0.0:1240/ws"
	addr := fmt.Sprintf("ws://%s/ws", t.ipaddress)

	// 建立连接
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		//log4.Fatal("dial:", err)
		t._FailHandler(t, "context fail")
		return
	}
	defer conn.Close()
	t._OnOpened(conn)
	byteBuffer := make([]byte, 0)
	bodyLen := 0
	// 循环读取数据
	for {
		// 读取客户端发送的数据
		_, buf, err := conn.ReadMessage()
		n := len(buf)
		//n, err := conn.Read(buf)
		if err != nil {
			//log.Printf("Error reading from connection: %v", err)
			t._OnClosed(conn, "Error reading from connection")
			return
		}
		if n == 0 {
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
				//t._connTime.Set(netId, time.Now())
				t._HandlerData(netId, msgId, content)
			}
		}

	}
}

func (t *WSTCPClient) SendMessage(msgId int32, msgData []byte) error {
	if t.conn == nil {
		return nil
	}
	//长度[msgid,content]
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

	// 发送响应消息
	err := t.conn.WriteMessage(websocket.BinaryMessage, msgPkg)
	//log.Printf("Error writing message: %s", err)
	return err
}

func (t *WSTCPClient) _OnOpened(conn *websocket.Conn) {
	netId := conn.RemoteAddr().String()

	t.conn = conn
	log4.Info("accept -ctx->netId  %s->%s", conn.LocalAddr().String(), netId)
	if t._ConnAddHandler != nil {
		t._ConnAddHandler(netId)
	}

}
func (t *WSTCPClient) _OnClosed(conn *websocket.Conn, prefix string) {
	netId := conn.RemoteAddr().String()
	t.conn = nil
	log4.Info("closed -ctx->netId  %s->%s", conn.LocalAddr().String(), netId)
	if t._ConnCloseHandler != nil {
		t._ConnCloseHandler(netId)
	}
}
func (t *WSTCPClient) Close() {
	if t.conn != nil {
		t.conn.Close()
		t.conn = nil
	}

}

func (t *WSTCPClient) UpdataTime() {

}

func (t *WSTCPClient) GetNetId() string {
	if t.conn != nil {
		return t.conn.RemoteAddr().String()
	}
	return "-1"
}
