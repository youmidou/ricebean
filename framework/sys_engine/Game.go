package sys_engine

import (
	"phoenix-tudou/framework/log4"
	"runtime/debug"
)

/*
	接口管理器

added by yh @ 2024/05/07 10:43
规则:
RegisterModule 注册模块
QueryModule 查询模块
*/
type IInterfaceMgr interface {
	RegisterModule(name *string, module *IModule)
	QueryModule(name string) *IModule
}

// ---------GameState----------------
type GameState int32 //游戏状态
const (
	GS_STOP     GameState = iota //停止
	GS_RUNNING                   //运行中
	GS_STOPPING                  //停止中
)

/*
	游戏基础类

added by yh @ 2024/01/20 10:43
	接口->IInterfaceMgr
规则:

	1.注册监听服务器
	2.注册监听客户端
*/

type Game struct {
	m_exit        bool //退出
	_gameState    GameState
	m_module_map  map[string]IModule //模块映射
	m_module_list []IModule          //模块列表
	logFileName   string             //日志文件名

	//m_inited_list  []IModule //已经初始化的模块
	//m_started_list []IModule //已经启动的模块
}

func NewGame() *Game {
	t := &Game{}
	t.m_module_map = make(map[string]IModule)
	t.m_module_list = make([]IModule, 0)
	//t.m_inited_list = make([]IModule, 0)
	//t.m_started_list = make([]IModule, 0)
	t._gameState = GS_STOP
	return t
}
func (t *Game) SetLogFileName(fileName string) {
	t.logFileName = fileName
}

func (t *Game) Start() {

}
func (t *Game) Run() {
	for _, module := range t.m_module_list {
		module.Init()
	}
	for name, module := range t.m_module_list {
		module.Start()
		t.m_module_list[name] = module
	}
	/*
		for name, module := range t.m_module_list {
			module.Init()
			//module.Update()
			module.Stop()
			module.Release()
			module.GetState()
		}
	*/
	t._gameState = GS_RUNNING

	go t._Loop()
}

func (t *Game) Stop() {
	t._gameState = GS_STOPPING
	for _, module := range t.m_module_list {
		module.Stop()
	}
	t._gameState = GS_STOP
}
func (t *Game) Dispose() {
	for _, module := range t.m_module_list {
		module.Dispose()
	}
}
func (t *Game) RegisterModule(name string, module IModule) {
	t.m_module_map[name] = module
	t.m_module_list = append(t.m_module_list, module)
}
func (t *Game) QueryModule(name string) IModule {
	m := t.m_module_map[name]
	if m != nil {
		return m
	}
	m = t.m_module_map[name]
	if m != nil {
		return m
	}
	m = t.m_module_map[name]
	if m != nil {
		return m
	}
	return nil
}
func (t *Game) GetWorkPath() string {

	return ""
}

// -------------protected-----------------------
func (t *Game) _Loop() {
	for {
		defer func() {
			if r := recover(); r != nil {
				err := debug.Stack()
				log4.Error("Game->Loop:%s %s", r, err)
			}
		}()
		for _, module := range t.m_module_list {
			if module != nil {
				module.Update()
			}
		}
	}
}

func (t *Game) SetMain(s string) {

}

func (t *Game) SetLog4Name(s string) {

}
