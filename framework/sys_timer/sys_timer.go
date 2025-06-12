package sys_timer

import (
	"context"
	"fmt"
	"phoenix-tudou/framework/log4"
	"runtime/debug"
	"time"
)

/*
		定时器
	  - author yh @ 2024-01-23 17:58
	    *注意:
*/
func NewTimer(interval time.Duration, loop int32, doTaskFunc func()) *Timer {
	t := &Timer{}
	t.interval = interval
	t.loop = loop
	t.loopIndex = loop
	t.doTaskFunc = doTaskFunc
	return t
}

type Timer struct {
	ctx        context.Context
	cancel     context.CancelFunc
	interval   time.Duration
	doTaskFunc func()
	_isRunning bool
	loop       int32
	loopIndex  int32
}

func (t *Timer) Start() {
	if t._isRunning {
		return
	}
	t._isRunning = true
	// 创建一个带有取消能力的上下文
	var ctx, cancel = context.WithCancel(context.Background())
	t.ctx = ctx
	t.cancel = cancel
	go t.doStartTimer(t.ctx, t.interval)
}
func (t *Timer) ReStart() {

}

func (t *Timer) Stop() {
	t._isRunning = false
	if t.cancel != nil {
		t.cancel()
	}
	t.ctx = nil
	t.cancel = nil
}
func (t *Timer) _doTaskFunc() {
	defer func() {
		if r := recover(); r != nil {
			err := debug.Stack()
			log4.Error("TimerManager->taskFunc:%s  %s", r, err)
			t.Stop()
		}
	}()
	t.doTaskFunc()
}
func (t *Timer) doStartTimer(ctx context.Context, interval time.Duration) {
	// 创建一个定时器，每隔一段时间触发一次
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			// 定时任务的逻辑
			//fmt.Println("xxxxTimer ticked at", time.Now())
			if t.loop != -1 {
				if t.loopIndex-1 < 0 {
					t.Stop()
					return
				}
				t.loopIndex--

			}
			go t._doTaskFunc()
		case <-ctx.Done():
			// 如果收到取消信号，结束定时器
			fmt.Println("Timer stopped.")
			return
		}
	}
}
