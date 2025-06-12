package sys_engine

type Session struct {
	handle     *func()
	rmi_module *IRMIModule
	netid      *NetID
}

func NewSession(handle *func(), rmi_module *IRMIModule, netid *NetID) *Session {
	t := &Session{}
	t.handle = handle
	t.rmi_module = rmi_module
	t.netid = netid
	return t

}

/*
	远程方法调用状态

added by yh @ 2024/05/11 12:11
*/
type IRMIModule interface {
	IModule
	StartServer(listen_port Port)
	Register(obj *RMIObject) bool
	/*
		创建会话
		@param netid 网络ID
		@param ip IP地址
		@param port 端口
		@return 会话
	*/
	CreateSession(netid *NetID, ip *IP, port *Port) *Session
	/*
		关闭会话
		@param session 会话
	*/
	CloseSession(session *Session)
	Call(session *Session, method *string, in_param []byte, out_param []byte) ERMIDispatchStatus
	/*
		启动异步线程
	*/
	StartAsyncThread(threadnum int)
}
