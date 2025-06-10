package mail

import "ricebean/examples/gameserver/lobby_svc/_interfaces"

type UserMailManager struct {
}

func NewUserMailManager() *UserMailManager {
	t := &UserMailManager{}
	return t
}
func (t *UserMailManager) SetLobbySvc(g _interfaces.LobbySvc) {

}
