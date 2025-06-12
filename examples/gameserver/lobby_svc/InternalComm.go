package lobby_svc

import (
	"context"
	"github.com/golang/protobuf/proto"
	"ricebean/examples/_def"
	pitaya "ricebean/pkg"
	"ricebean/z_tools/ProtoToCS/Protocal/pbs"
)

const (
	GatewaySvcRoute    = "Gateway.GatewaySvc.OnReceive"
	GlobalSvcRoute     = "Global.GlobalSvc.OnReceive"
	DataAccessSvcRoute = "DataAccess.DataAccessSvc.OnReceive"
	GameWorldSvcRoute  = "GameWorld.GameWorldSvc.OnReceive"
)

// Impl
type InternalComm struct {
	app pitaya.Pitaya
}

// 内部通讯
func Instance() *InternalComm {
	return &InternalComm{}
}
func (t *InternalComm) SetApp(app pitaya.Pitaya) {
	t.app = app
}

func (t *InternalComm) RPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message) error {
	//ctx := context.Background()
	return t.app.RPC(ctx, routeStr, reply, arg)
}
func (t *InternalComm) RPCTo(ctx context.Context, serverID, routeStr string, reply proto.Message, arg proto.Message) error {

	return t.app.RPCTo(ctx, serverID, routeStr, reply, arg)
}

func (t *InternalComm) SendRPCGateway(netId _def.GSNetID, userId _def.UserID, msgId int32, msgData proto.Message) {
	ctx := context.Background()
	//serverId := pbs.ServerId_Gateway.String()
	//routeStr := pbs.Cmd_Gateway_GatewaySvc_OnGatewayReceive.String()
	msgBytes, err := proto.Marshal(msgData)
	if err != nil {
		return
	}
	req := &pbs.Net_Internal_ForwardSvcGeneralRPCReq{}
	req.NetId = string(netId)
	req.UserId = string(userId)
	req.MsgType = 1
	req.MsgId = msgId
	req.MsgData = msgBytes
	ret := &pbs.Net_Internal_ForwardSvcGeneralRPCRet{}
	t.app.RPC(ctx, "Gateway.GatewaySvc.OnReceive", ret, req)
}

// InternalComm
func (t *InternalComm) NetDisconnect(netId _def.GSNetID, userId _def.UserID) {
	ctx := context.Background()
	req := &pbs.Net_Internal_ForwardSvcGeneralRPCReq{}
	req.NetId = string(netId)
	req.UserId = string(userId)
	req.MsgType = 2
	ret := &pbs.Net_Internal_ForwardSvcGeneralRPCRet{}
	str := []string{string(userId)}
	t.app.SendKickToUsers(str, "frontendType")
	t.app.RPC(ctx, "Gateway.GatewaySvc.OnReceive", ret, req)
}
func (t *InternalComm) SendRPCGameWorld(netId _def.GSNetID, userId _def.UserID, msgId int32, msgData proto.Message) {
	ctx := context.Background()
	//serverId := pbs.ServerId_Gateway.String()
	//routeStr := pbs.Cmd_Gateway_GatewaySvc_OnGatewayReceive.String()
	msgBytes, err := proto.Marshal(msgData)
	if err != nil {
		return
	}
	req := &pbs.Net_Internal_ForwardSvcGeneralRPCReq{}
	req.NetId = string(netId)
	req.UserId = string(userId)
	req.MsgType = 1
	req.MsgId = msgId
	req.MsgData = msgBytes
	ret := &pbs.Net_Internal_ForwardSvcGeneralRPCRet{}
	t.app.RPC(ctx, "Gateway.GatewaySvc.OnReceive", ret, req)
}
