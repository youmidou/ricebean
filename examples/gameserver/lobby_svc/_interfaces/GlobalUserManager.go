package _interfaces

import "ricebean/examples/_def"

type GlobalUser interface {
	GetName() string
}
type GlobalUserManager interface {
	AddUser(platName _def.PlatName, userId _def.UserID, role_name _def.GameName, netId _def.GSNetID) bool
	RemoveUser(userId _def.UserID) bool
	RemoveUserOnGateway(gateway_netid _def.GSNetID) int32

	GetUser(userId _def.UserID) GlobalUser
	GetUserByNetID(netId _def.GSNetID) GlobalUser
	GetUserByName(name _def.GameName) GlobalUser

	GetUserCount() int
}
