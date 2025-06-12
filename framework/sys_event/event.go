/*
* 事件调度器
added by yh @ 2023/10/23 18:12
注意:
*/
package sys_event

import (
	"phoenix-tudou/framework/log4"
)

type EventDispatcher struct {
	cacheEvents map[string]func(interface{})
}

func (t *EventDispatcher) Init() {
	t.cacheEvents = make(map[string]func(interface{}))
}
func (t *EventDispatcher) AddListener(cmd string, fun func(interface{})) {
	if cmd == "" {
		log4.Error("cmd need to pass value!")
		return
	}
	if fun == nil {
		log4.Error("fun need to pass value!")
		return
	}
	if t.cacheEvents[cmd] == nil {
		t.cacheEvents[cmd] = fun
	} else {
		log4.Error("cmd:%s fun:%s", cmd, fun)
	}
}

// 派发
func (t *EventDispatcher) Broadcast(cmd string, args interface{}) {
	_fun := t.cacheEvents[cmd]
	if _fun != nil {
		_fun(args)
	}
}

// 卸载
func (t *EventDispatcher) RemoveListener(cmd string) {
	delete(t.cacheEvents, cmd)
}
func (t *EventDispatcher) Dispose() {
	t.cacheEvents = nil
}
