/*
*
added by yh @ 2023/5/19 10:01
注意:
*/
package sys_rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/z_Tools/ProtoToCS/Protocal/pb"
)

type GateToGameServer struct {
	pb.UnimplementedGateToGameServer
	Num int
}

// 监听 ":1234"
func (t *GateToGameServer) Listen(addrPort string) {
	lis, err := net.Listen("tcp", addrPort)
	if err != nil {
		log4.Panic(fmt.Sprintf("无法监听端口: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterGateToGameServer(s, t) //&GateToGameServer{}
	log4.Info("服务启动于%s", addrPort)
	if err := s.Serve(lis); err != nil {
		log4.Error("无法提供服务: %s", err.Error())
	}
}
func (t *GateToGameServer) ProcessGateToGame(ctx context.Context, req *pb.GateToGameReq) (*pb.GateToGameRet, error) {
	log4.Info("req=" + req.ReqId)

	//err := errors.New("爆个错耍耍==================")
	return &pb.GateToGameRet{RetId: "GameServer->Hello****"}, nil
}
