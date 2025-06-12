package sys_map

/*
	LinkedHashMap

added by yh @ 2024-01-19 21:11
注意：参照Java LinkedHashMap 实现有序map
使用LinkedHashMap：LinkedHashMap是HashMap的子类，
它维护了一个双向链表来维护插入顺序或者访问顺序，
因此能够以插入顺序或者访问顺序遍历元素。
LinkedHashMap的性能介于HashMap和TreeMap之间。
*/
type LinkedHashMap struct {
	_HashMap map[string]any
}

func NewLinkedHashMap[K string, V any]() {
	t := &LinkedHashMap{}
	t._HashMap = make(map[string]any)
}
func (t *LinkedHashMap) Get(key string) {
	//xx := max[int64](100, 1000)
	//if xx == 1 {
	//}
}
func (t *LinkedHashMap) Set(key string, v any) {

}
func (t *LinkedHashMap) Delete(key string) {

}

/*
	节点

added by yh @ 2024-01-19 21:11
*/
type Node struct {
	//-------双向链表头节点----------------------------
	before string //前
	after  string //后

}

// 下一个指针
func (t *Node) Next() any {

	return t.after
}

func (t *Node) Previous() any {
	return t.after
}

// 第一个
func (t *Node) First() any {
	return t.after
}

// 最后一个
func (t *Node) Last() any {
	return t.after
}
