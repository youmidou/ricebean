package lobby_svc

import (
	"github.com/golang/protobuf/proto"
	"phoenix-tudou/examples/_def"
	"phoenix-tudou/framework/sys_net"
	"phoenix-tudou/z_Tools/ProtoToCS/Protocal/pb"
)

type LobbySvcMsgHandler struct {
	m_msg_handler_list map[int32]func(string, *GlobalUser, int32, []byte) (sys_net.IRequestHandler, error) //proto.Message
	LobbySvc           *LobbySvc
}

// 消息处理 Register
func NewLobbySvcMsgHandler(LobbySvc *LobbySvc) *LobbySvcMsgHandler {
	t := &LobbySvcMsgHandler{}
	t.LobbySvc = LobbySvc
	t.m_msg_handler_list = make(map[int32]func(string, *GlobalUser, int32, []byte) (sys_net.IRequestHandler, error))
	t.register()
	return t
}

func (t *LobbySvcMsgHandler) register() {
	t.m_msg_handler_list[int32(pb.Cmd_Lobby_AuthorizationLogin)] = t.AuthorizationLoginHandler
	t.m_msg_handler_list[int32(pb.Cmd_Lobby_AuthorizationEnterGame)] = t.EnterGameHandler
}

func (t *LobbySvcMsgHandler) Process(netId string, globalUser *GlobalUser, msgId int32, bytes []byte) (sys_net.IRequestHandler, error) {
	if t.m_msg_handler_list[msgId] != nil {
		return t.m_msg_handler_list[msgId](netId, globalUser, msgId, bytes)
	} else {

	}
	return nil, nil
}

// OnEnterLobbySvcReq
func (t *LobbySvcMsgHandler) AuthorizationLoginHandler(netId string, globalUser *GlobalUser, msgId int32, bytes []byte) (sys_net.IRequestHandler, error) {
	req := &pb.Net_Lobby_AuthorizationLoginReq{}
	//ret := &pb.Net_Lobby_AuthorizationLoginRet{}
	proto.Unmarshal(bytes, req)
	//token := req.Token
	//token.key
	//token.userId
	userId := "1111"
	pname := "pname"
	role_name := "role_name"
	if "key" != "token.key" {
		//_def.GameName(role_name), "scene_id", "SessionKeyError"
		Instance().NetDisconnect(_def.GSNetID(netId), _def.UserID(userId))
	}
	//	if ((unsigned int)EngineAdapter::Instance().Time() > (unsigned int)egsr->time + m_max_login_interval)
	if "key" != "token.key" {
		//EngineAdapter.Instance().NetDisconnect(_def.GSNetID(netId), userId, _def.GameName(role_name), "scene_id", "SessionKeyError")
	}
	isok := t.LobbySvc.GetUserManager().AddUser(_def.PlatName(pname), _def.UserID(userId), _def.GameName(role_name), _def.GSNetID(netId))
	if isok {
		//
	}
	user := t.LobbySvc.GetUserManager().GetUser(userId)
	//user := t.LobbySvc.GetUserManager().GetUserByNetID(netId)
	if user == nil {

	}
	//EngineAdapter.Instance().NetSend(_def.GSNetID(netId), &pb.Net_Login_LoginReq{})
	return nil, nil
}

func (t *LobbySvcMsgHandler) EnterGameHandler(netId string, globalUser *GlobalUser, msgId int32, bytes []byte) (sys_net.IRequestHandler, error) {
	return nil, nil
}
