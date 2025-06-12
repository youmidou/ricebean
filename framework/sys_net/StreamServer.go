package sys_net

/*
	permanent server

added by yh @ 2023-04-24
*/
func StreamServerWS(ipaddress string, HandlerData func(netId string, msgId int32, content []byte),
	FailHandler func(t IServer, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string),
) *WSTCPServer {
	var server = NewWSTCPServer(ipaddress,
		HandlerData,
		FailHandler,
		ConnAddHandler,
		ConnCloseHandler,
	)
	//ipaddress:=fmt.Sprintf( "%s:%d",host_ip,host_port)
	return server
}

/*
func StreamServer(ipaddress string, maxConn int32, handler func(conn net.Conn)) *TCPServer {
	var server = NewTCPServer(ipaddress, maxConn, "sss", handler)
	//ipaddress:=fmt.Sprintf( "%s:%d",host_ip,host_port)
	return server
}
*/

/*
func StreamClient(ipaddress string,

	handlerData func(netId string, msgId int32, content []byte),
	failHandler func(t ITCPServer, prefix string),
	ConnAddHandler func(netId string),
	ConnCloseHandler func(netId string),

	) *TCPServer {
		//var client = nil
		return nil
	}
*/
