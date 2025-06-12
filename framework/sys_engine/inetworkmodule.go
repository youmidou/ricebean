package sys_engine

const NETWORK_MODULE = "NetworkModule"

type NetID string //网络ID
type IP string    //IP地址
type Port uint16  //端口

type IEngineNetCallback interface {
	OnAccept(listen_port Port, netid NetID, ip IP, port Port) //接受连接
	OnRecv(netid NetID, data []byte)                          //接收数据
	OnDisconnect(netid NetID)                                 //断开连接
	OnConnect(netid NetID)                                    //连接
}
type INetworkModule interface {
	IModule
	Init()
	Start()
	Update()
	Stop()
	Dispose()
	RegisterCallback(callback *IEngineNetCallback)
	/*
		注册监听服务器
		@param port 端口
		@param backlog 最大连接数
		@param netid_out 网络ID
		@param ip_bind 绑定的IP
		@return 是否成功
	*/
	Listen(port Port, backlog int, netid_out *NetID, ip_bind *byte) bool
	/*
		连接服务器
		@param ip IP地址
		@param port 端口
		@param netid_out 网络ID
		@param time_out 超时时间 默认3000ms
	*/
	Connect(ip IP, port Port, netid_out *NetID, time_out uint64) bool
	ConnectAsyn()
	Send()
	/*
		断开连接
		IEngineNetCallback.OnDisconnect
		@param netid 网络ID
	*/
	Disconnect(netid NetID)
}
