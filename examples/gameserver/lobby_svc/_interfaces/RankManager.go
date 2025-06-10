package _interfaces

type RankManager interface {
	OnUserLogin(user GlobalUser)
	OnUserLogout(user GlobalUser)
}
