package sys_engine

/*
	API接口模块接口

added by yh @ 2024/05/11 11:24
注意:
*/
type IAPIModule interface {
	IModule
	Init()
	Start()
	Update()
	Stop()
	Release()
	Free()
}
