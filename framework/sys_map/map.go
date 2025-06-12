package sys_map

import (
	"unsafe"
)

/*
	NormalMap 是一个泛型并发安全的 Map

added by yh @ 2023/5/17 11:35
列子: t._connMap = sys_map.NewNormalMap[string, *net.Conn]()
注意:辉哥提醒 sync.Map 并不会对存储在 map 中的对象进行额外的保护
*/

// NormalMap 是一个泛型并发安全的 Map
type NormalMap[K comparable, V any] struct {
	m      map[K]V // 存储键值对的 map
	length int64   // 使用 int64 来存储键值对数量
}

// NewNormalMap 创建一个新的 NormalMap 实例
func NewNormalMap[K comparable, V any]() *NormalMap[K, V] {
	t := &NormalMap[K, V]{
		m: make(map[K]V),
	}
	return t
}

// Load 根据 key 获取值
func (c *NormalMap[K, V]) Get(key K) V {
	return c.m[key]
}

// Store 设置 key 对应的值，如果 key 已存在则更新值
func (c *NormalMap[K, V]) Set(key K, value V) {
	c.m[key] = value
}

// Delete 删除 key 对应的值
func (c *NormalMap[K, V]) Delete(key K) {
	delete(c.m, key)
}

// Range 遍历所有的 key-value 对
func (c *NormalMap[K, V]) Range(f func(key K, value V) bool) {
	for k, v := range c.m {
		f(k, v)
	}
}

// Len 获取当前存储的键值对数量
func (c *NormalMap[K, V]) Len() int {
	// 返回计数器的值，并将 int64 转换为 int
	return len(c.m)
	//return int(atomic.LoadInt64(&c.length))
}

// Clear 清理所有键值对
func (c *NormalMap[K, V]) Clear() {
	c.m = make(map[K]V)
}

// Sizeof 当前并发数组内存大小
func (c *NormalMap[K, V]) Sizeof() int64 {
	size := unsafe.Sizeof(c.m)
	return int64(size)
}
