package EventDispatcher

import (
	"fmt"
	"reflect"
	"ricebean/framework/log4"
	"sync"
)

var instance *EventDispatcher
var once sync.Once

func GetInstance() *EventDispatcher {
	once.Do(func() {
		instance = NewEventDispatcher()
	})
	return instance
}

type EventDispatcher struct {
	callBackMap map[string]interface{}
}

func NewEventDispatcher() *EventDispatcher {
	t := &EventDispatcher{}
	t.callBackMap = make(map[string]interface{})
	return t
}

// RegisterCallBack
func (t *EventDispatcher) AddListener(key string, callBack interface{}) {
	t.callBackMap[key] = callBack
}
func (t *EventDispatcher) Broadcast(key string, args ...interface{}) []interface{} {
	if callBack, ok := t.callBackMap[key]; ok {
		in := make([]reflect.Value, len(args))
		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}
		outList := reflect.ValueOf(callBack).Call(in)
		result := make([]interface{}, len(outList))
		for i, out := range outList {
			result[i] = out.Interface()
		}
		return result
	} else {
		log4.Panic(fmt.Errorf("callBack(%s) not found", key).Error())
	}
	return nil
}
func (t *EventDispatcher) RemoveListener(key string) {
	t.callBackMap[key] = nil
}
