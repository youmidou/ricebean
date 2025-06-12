package sys_net

import (
	"net"
	"time"
)

/*
	EventServer eventHandler EventHandler

added by yh @ 2024-03-11
*/
type EventServer interface {
	//OnInitComplete(server *TCPServer) //初始化连接
	OnShutdown()                 //关闭时候
	OnOpened(c net.Conn)         //有连接进来
	OnClosed(c net.Conn)         //断开连接时候
	React([]byte, net.Conn)      //接收消息
	Tick() (delay time.Duration) //action
}
