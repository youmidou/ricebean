package _interfaces

type Activity interface {
}
type ActivityManager interface {
	IsActivityOpen(activity_type int) bool
	IsActivityClose(activity_type int) bool

	OnUserLogin(user GlobalUser)
	OnUserLogout(user GlobalUser)
}
