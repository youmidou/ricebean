package _user

import "phoenix-tudou/examples/_def"

type GlobalUser struct {
	netId             _def.GSNetID
	userId            _def.UserID
	m_scene_id        _def.SceneIndex
	platName          _def.PlatName
	gameName          _def.GameName
	globalUserManager *GlobalUserManager
}

func NewGlobalUser(platName _def.PlatName, userId _def.UserID, gameName _def.GameName, netId _def.GSNetID, globalUserManager *GlobalUserManager) *GlobalUser {
	t := &GlobalUser{}
	t.platName = platName
	t.userId = userId
	t.gameName = gameName
	t.netId = netId
	t.globalUserManager = globalUserManager
	return t
}
func (t *GlobalUser) OnSyncRoleBaseInfo() {

}
func (t *GlobalUser) SetSceneId(scene_index _def.SceneIndex) {
	t.m_scene_id = scene_index
}
func (t *GlobalUser) GetSceneID() _def.SceneIndex {
	return t.m_scene_id
}
func (t *GlobalUser) GetPlatName() _def.PlatName {
	return t.platName
}
func (t *GlobalUser) GetName() string {
	return string(t.gameName)
}
