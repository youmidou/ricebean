package sys_engine

type IGameModule interface {
	IModule
	StopGame()
	GetWorkPath(path *string) bool
}
