package sys_engine

const STD_COMMAND_MODULE = "StdCommandModule"

type ICmdCallback interface {
	OnCmd(cmd string)
}
type IStdCommandModule interface {
	IModule
	RegisterCallback(callback *ICmdCallback)
}
