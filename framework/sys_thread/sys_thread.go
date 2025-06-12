package sys_thread

import (
	"context"
	"fmt"
	"phoenix-tudou/framework/log4"
	"runtime/debug"
)

/*
	创建一条可控制的线程

author yh @ 2024-01-24 15:09

	注意:本线程可以外部杀掉
*/
func NewThread(doTaskFunc func()) *Thread {
	t := &Thread{}
	// 创建一个带有取消能力的上下文
	ctx, cancel := context.WithCancel(context.Background())
	t.ctx = ctx
	t.cancel = cancel
	t.doTaskFunc = doTaskFunc
	return t
}

type Thread struct {
	ctx        context.Context
	cancel     context.CancelFunc
	doTaskFunc func()
	_isRunning bool
}

func (t *Thread) Start() {
	if t._isRunning {
		return
	}
	t._isRunning = true
	go t.doAThread(t.ctx)

}
func (t *Thread) Stop() {
	t._isRunning = false
	if t.cancel != nil {
		t.cancel()
	}

}
func (t *Thread) _doTaskFunc() {
	defer func() {
		if r := recover(); r != nil {
			err := debug.Stack()
			log4.Error("Thread->doTaskFunc 报错:%s  %s", r, err)
			t.Stop()
		}
	}()
	t.doTaskFunc()
}
func (t *Thread) doAThread(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// 当上下文被取消时，退出 goroutine
			fmt.Println("Goroutine cancelled.")
			return
		default:
			// 执行一些任务
			fmt.Println("Working...")
			t._doTaskFunc()
			t.Stop()
		}
	}

}
