package activity

import (
	"ricebean/examples/GameServer/lobby_svc/_interfaces"
)

type ActivityManager struct {
}

func (m ActivityManager) IsActivityOpen(activity_type int) bool {
	//TODO implement me
	return false
}

func (m ActivityManager) IsActivityClose(activity_type int) bool {
	return false
}

func (m ActivityManager) OnUserLogin(user _interfaces.GlobalUser) {

}

func (m ActivityManager) OnUserLogout(user _interfaces.GlobalUser) {

}

func (m ActivityManager) SetLobbySvc(svc _interfaces.LobbySvc) {

}

func NewActivityManager() *ActivityManager {
	t := &ActivityManager{}
	return t
}
