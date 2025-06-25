package lobby_svc

import (
	"context"
	"github.com/golang/protobuf/proto"
	_interfaces2 "ricebean/examples/gameserver/lobby_svc/_interfaces"
	"ricebean/examples/gameserver/lobby_svc/sys_modules/activity"
	"ricebean/examples/gameserver/lobby_svc/sys_modules/mail"
	"ricebean/framework/log4"

	"ricebean/framework/sys_net"
	pitaya "ricebean/pkg"
	"ricebean/pkg/component"

	"ricebean/z_tools/ProtoToCS/Protocal/pbs"
)

const (
	MT_NET_RECV_MSG          = 1
	MT_NET_DISCONNECT_NOTICE = 2
)

/*
//modules.Base
*/
func NewLobbySvc(app pitaya.Pitaya) *LobbySvc {
	return &LobbySvc{app: app}
}

type LobbySvc struct {
	component.Base
	app                pitaya.Pitaya
	globalUserManager  *GlobalUserManager
	m_usermail_manager *mail.UserMailManager
	m_activity_manager *activity.ActivityManager
	msgHandler         *LobbySvcMsgHandler
	_request_map       map[int32]func(msgId int32, bytes []byte) (sys_net.IRequestHandler, error)
}

func (t *LobbySvc) Init() {
	t._request_map = make(map[int32]func(msgId int32, bytes []byte) (sys_net.IRequestHandler, error))
	//全局用户
	t.globalUserManager = NewGlobalUserManager()

	t.m_usermail_manager = mail.NewUserMailManager()
	t.m_activity_manager = activity.NewActivityManager()

	//-------注册接收消息--------------------
	t.msgHandler = NewLobbySvcMsgHandler(t)

	//t.RegisterHandleMessage(pb.Cmd_Login_Login, t.LoginHandle)
	//InternalCommInstance().SetApp(t.app)
	//t.m_user_manager.SetLobbySvc(t)
	//t.m_activity_manager.SetLobbySvc(t)

}

// 内部接收消息
func (t *LobbySvc) InternalReceivingMessage() {
	//InternalMessageProcessing

}

// 接收网关数据
func (t *LobbySvc) OnGatewayReceive(ctx context.Context, req *pbs.Net_InternalMessagePacket) (*pbs.Net_InternalMessagePacket, error) {

	ret := &pbs.Net_InternalMessagePacket{}
	switch req.MsgType {
	case int32(pbs.PacketType_Net_ReceiveMsg): //收到消息
		ret, _ = t.RequestProcessingMessage(req)
	case int32(pbs.PacketType_Net_Disconnect_Notice): //断开通知
		//处理断开
		t.OnUserDisconnect(ret)
		return nil, nil
		//InternalComm.Instance().NetDisconnect(_def.GSNetID(req.NetId), _def.UserID(req.UserId))
	}
	/*
		//sys_net.MessageContentUnpack()
		//req.ReqId
		//t.app.SendPushToUsers()
		session := t.app.GetSessionFromCtx(ctx)
		if session == nil {
			return nil, nil
		}
		session.Bind(ctx, "12345")
		session.GetHandshakeData()
		netId := session.ID()
		//userId := session.UID()

		user := t.m_user_manager.GetUserByNetID(netId)
		if user == nil {
			session.Close()
			return nil, nil
		}*/

	return ret, nil
}

// 注册处理消息
func (t *LobbySvc) RegisterHandleMessage() {
	//t.app.
}

func (t *LobbySvc) RequestProcessingMessage(req *pbs.Net_InternalMessagePacket) (*pbs.Net_InternalMessagePacket, error) {
	/*
		user := t.m_globaluser_manager.GetUserByNetID(req.NetId)
		switch req.MsgId {
		case int32(pb.Cmd_Login_Login):
		default:
			if user == nil {
				//ctxGate.RemoveConnect(ctx.Cid)
				//log4.Info("发现User==nil gate=%s reqId=%d", gateReq.GateCid, gateReq.ReqId)
				time.Sleep(3 * time.Second)
				//ctxGate.Send(int32(pbs.Cmd_Central_GateKillConnection), killBytes)
				return nil, nil
			}
		}
		if req.MsgId == 0 {
			return nil, nil
		}
		moduleId := sys_base.GetCmdModule(req.MsgId)
		switch moduleId {
		case pb.CmdModule_Lobby: //大厅消息
			resp, _ := t.msgHandler.Process(req.NetId, user, req.MsgId, req.MsgData)
			if resp != nil {
				r := resp.GetResponse()
				msgData, _ := proto.Marshal(r.Ret)
				t.ForwardRPCGateway(req.NetId, req.UserId, r.RetId, msgData)
			}
			return nil, nil
		default:
			log4.Error("GateToCentral unknown message pending gateReq.ReqId = %s", req.MsgId)
			return nil, nil
		}
	*/
	ret := &pbs.Net_InternalMessagePacket{}
	//resp := RequestProcessor.Process(ctx, gateReq.ReqId, gateReq.ReqBytes)
	return ret, nil
}

// 全局玩家管理器
func (t *LobbySvc) GetUserManager() *GlobalUserManager {
	return t.globalUserManager
}

// 全局玩家管理器
func (t *LobbySvc) GetUserMailManager() *mail.UserMailManager {
	return t.m_usermail_manager
}

// 处理用户断开连接
func (t *LobbySvc) OnUserDisconnect(ret *pbs.Net_InternalMessagePacket) {

}

func (t *LobbySvc) SendRPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message) error {
	return t.app.RPC(ctx, routeStr, reply, arg)
}

// 转发网关
func (t *LobbySvc) ForwardRPCGateway(netId string, userId string, msgId int32, msgData []byte) bool {
	ctx := context.Background()
	req := &pbs.Net_Internal_GatewayRPCGameWorldReq{}
	req.NetId = netId
	req.UserId = userId
	req.MsgId = msgId
	req.MsgData = msgData
	ret := &pbs.Net_Internal_GatewayRPCGameWorldRet{}
	err := t.app.RPC(ctx, "gateway.onuserrecv", ret, req)
	if err != nil {
		log4.Error("ForwardRPCGateway [%s,%s,%s] fail ...", netId, userId, msgId)
		return false
	}
	return true
}
func (t *LobbySvc) SendToUserGameWorld() {

}

// 转发游戏世界
func (t *LobbySvc) ForwardRPCGameWorld(netId string, userId string, msgId int32, msgData []byte) *pbs.Net_Internal_GlobalRPCGameWorldRet {
	ctx := context.Background()
	req := &pbs.Net_Internal_GlobalRPCGameWorldReq{}
	req.NetId = netId
	req.UserId = userId
	req.MsgId = msgId
	req.MsgData = msgData
	ret := &pbs.Net_Internal_GlobalRPCGameWorldRet{}
	err := t.app.RPC(ctx, "gameworld.onuserrecv", ret, req)
	if err != nil {
		log4.Error("转发->gameworld.onuserrecv fail")
		return nil
	}
	return ret
}

// 转发dataaccess msgData proto.Message
func (t *LobbySvc) ForwardRPCDataAccess(netId string, userId string, msgId int32, msgData []byte) proto.Message {
	ctx := context.Background()
	req := &pbs.Net_Internal_GlobalRPCDataAccessReq{}
	req.NetId = netId
	req.UserId = userId
	req.MsgId = msgId
	req.MsgData = msgData

	ret := &pbs.Net_Internal_GlobalRPCDataAccessRet{}
	err := t.app.RPC(ctx, "dataaccess.onuserrecv", ret, req)
	if err != nil {
		log4.Error("转发->dataaccess.onuserrecv fail")
		return nil
	}
	return ret
}

func (t *LobbySvc) GetActivityManager() _interfaces2.ActivityManager {
	//TODO implement me
	return t.m_activity_manager
}

func (t *LobbySvc) GetRankManager() _interfaces2.RankManager {
	return nil
}
