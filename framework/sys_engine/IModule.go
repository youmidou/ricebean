package sys_engine

const (
	Succeed = iota
	Fail
	Pending
)

// --------State-------------------
type ModuleState int32 //模块状态
const (
	ST_Initing   ModuleState = iota //初始化中
	ST_Inited                       //初始化完成
	ST_Starting                     //启动中
	ST_Started                      //启动完成
	ST_Updating                     //更新中
	ST_Updated                      //更新完成
	ST_Stopping                     //停止中
	ST_Stopped                      //停止完成
	ST_Releasing                    //释放中
	ST_Released                     //释放完成
	ST_Fail                         //失败
)

/*
	功能模块

added by yh @ 2024/05/06 18:46
规则:

	1.注册监听服务器
	2.注册监听客户端
*/
type IModule interface {
	GetState() ModuleState //获取状态
	Init()                 //初始化
	Update()               //更新
	Start()                //开始
	Stop()                 //停止
	Dispose()              //析构函数函数

}
type Module struct {
	m_interface_mgr IInterfaceMgr
	State           ModuleState
}

func NewModule() *Module {
	t := &Module{}
	return t
}
func (t *Module) Init() {
	t.State = ST_Initing
	t.State = ST_Inited
}

func (t *Module) Update() {
	//t.State = ST_Updating
	//t.State = ST_Updated
}

func (t *Module) Start() {
	t.State = ST_Starting
	t.State = ST_Started
}
func (t *Module) Stop() {
	t.State = ST_Stopping
	t.State = ST_Stopped
}
func (t *Module) Dispose() {
	t.State = ST_Releasing
	t.State = ST_Released
}
func (t *Module) GetState() ModuleState {
	return t.State
}

/*
func (t *Module) Release() {

}
func (t *Module) Free() {

}
*/

func (t *Module) Interface() IInterfaceMgr {
	return t.m_interface_mgr
}
