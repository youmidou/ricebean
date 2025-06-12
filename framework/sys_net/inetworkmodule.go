package sys_net

type NetID string //网络ID
type IP string    //IP
type Port int16   //端口

type IEngineNetCallback interface {
	OnConnect(netId NetID)                 //连接成功
	OnAccept(listenPort Port, netId NetID) //接受连接
	OnRecv(netId NetID, data []byte)       //接受数据
	OnDisconnect(netId NetID)              //断开连接

}
