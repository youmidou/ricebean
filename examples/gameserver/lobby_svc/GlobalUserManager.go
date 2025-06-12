package lobby_svc

import (
	"ricebean/examples/_def"
	"ricebean/examples/gameserver/lobby_svc/_interfaces"
)

type GlobalUserManager struct {
	m_user_list     map[_def.UserID]*GlobalUser
	m_net_user_map  map[_def.GSNetID]*GlobalUser
	m_name_user_map map[_def.GameName]*GlobalUser
	m_global_server _interfaces.LobbySvc
}

func NewGlobalUserManager() *GlobalUserManager {
	return &GlobalUserManager{
		m_user_list:     make(map[_def.UserID]*GlobalUser),
		m_net_user_map:  make(map[_def.GSNetID]*GlobalUser),
		m_name_user_map: make(map[_def.GameName]*GlobalUser),
	}
}
func (t *GlobalUserManager) Init() {

}

func (t *GlobalUserManager) AddUser(platName _def.PlatName, userId _def.UserID, role_name _def.GameName, netId _def.GSNetID) bool {
	user := NewGlobalUser(platName, userId, role_name, netId, t)
	t.m_user_list[userId] = user
	t.m_net_user_map[netId] = user
	t.m_name_user_map[role_name] = user
	//t.server.SendRPC()

	t.m_global_server.GetActivityManager().OnUserLogin(user)
	t.m_global_server.GetRankManager().OnUserLogin(user)

	return true
}

func (t *GlobalUserManager) SetLobbySvc(server _interfaces.LobbySvc) {
	t.m_global_server = server
}
func (t *GlobalUserManager) GetUser(userId string) *GlobalUser {
	return nil
}
func (t *GlobalUserManager) GetUserByNetID(netId string) *GlobalUser {
	return nil
}
