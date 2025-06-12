package sys_net

import (
	"github.com/panjf2000/gnet"
	"log"
	"sync/atomic"
	"time"
)

/*
	GNetTCPServer

added by yh @ 2024/02/28 15:03
注意:
*/
type GNetTCPServer struct {
	*gnet.EventServer
	connections int64 // 连接计数器
}

func NewGNetTCPServer() *GNetTCPServer {
	t := &GNetTCPServer{}
	err := gnet.Serve(t, "tcp://0.0.0.0:1250", gnet.WithMulticore(true))
	if err != nil {
		log.Fatalf("Server start error: %v\n", err)
	}
	return t
}

// 初始化完成时
func (es *GNetTCPServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Server is listening on %s\n", srv.Addr.String())
	return
}

// 关机时
func (es *GNetTCPServer) OnShutdown(svr gnet.Server) {

}

func (es *GNetTCPServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	// 新连接建立时增加连接计数器
	atomic.AddInt64(&es.connections, 1)
	return
}

func (es *GNetTCPServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	// 连接断开时减少连接计数器
	atomic.AddInt64(&es.connections, -1)
	return
}
func (es *GNetTCPServer) PreWrite(c gnet.Conn) {

}
func (es *GNetTCPServer) AfterWrite(c gnet.Conn, b []byte) {

}

// 接收消息
func (es *GNetTCPServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	// 响应消息，这里简单地将收到的消息原样返回
	out = frame
	//es.NumEventLoopConnections(c)
	// 获取当前在线连接数
	log.Printf("Current number of connections: %d\n", atomic.LoadInt64(&es.connections))
	// 异步发送消息
	c.AsyncWrite(out)
	return
}
func (es *GNetTCPServer) Tick() (delay time.Duration, action gnet.Action) {
	// 定时任务：每隔一秒钟打印一次日志
	log.Println("Tick!")

	// 返回下一次定时任务执行的间隔时间
	return time.Second, gnet.None

}
