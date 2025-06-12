package gevent
/*
gevent 线程管理器
author:yh 2023.04.24
*/
//小绿叶
type Greenlet struct {
	done bool
}
//产卵 创建一个线程thread
func (t *Greenlet)Spawn(linecall func())  {

	//done := make(chan bool)

}