package sys_map

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

/*
	ConcurrentMap 是一个泛型并发安全的 Map

added by yh @ 2023/5/17 11:35
列子: t._connMap = sys_map.NewConcurrentMap[string, *net.Conn]()
注意:辉哥提醒 sync.Map 并不会对存储在 map 中的对象进行额外的保护
*/

// ConcurrentMap 是一个泛型并发安全的 Map
type ConcurrentMap[K comparable, V any] struct {
	m      sync.Map // 内部使用 sync.Map 来存储键值对
	length int64    // 使用 int64 来存储键值对数量
}

// NewConcurrentMap 创建一个新的 ConcurrentMap 实例
func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{}
}

// Load 根据 key 获取值
func (c *ConcurrentMap[K, V]) Get(key K) V {
	value, ok := c.m.Load(key)
	if ok {
		return value.(V) //, true
	}
	var zero V
	return zero //, false
}

// Store 设置 key 对应的值，如果 key 已存在则更新值
func (c *ConcurrentMap[K, V]) Set(key K, value V) {
	_, loaded := c.m.LoadOrStore(key, value)
	if !loaded {
		// 如果 key 是新添加的，则递增计数器
		atomic.AddInt64(&c.length, 1)
	} else {
		// 如果 key 已存在，则更新值
		c.m.Store(key, value)
	}
}

// LoadOrStore 获取 key 对应的值，若不存在则设置并返回新值
func (c *ConcurrentMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := c.m.LoadOrStore(key, value)
	if !loaded {
		// 如果 key 是新添加的，则递增计数器
		atomic.AddInt64(&c.length, 1)
	}
	return actual.(V), loaded
}

// Delete 删除 key 对应的值
func (c *ConcurrentMap[K, V]) Delete(key K) {
	_, loaded := c.m.Load(key)
	if loaded {
		// 如果 key 存在，则递减计数器
		c.m.Delete(key)
		atomic.AddInt64(&c.length, -1)
	}
}

// Range 遍历所有的 key-value 对
func (c *ConcurrentMap[K, V]) Range(f func(key K, value V) bool) {
	c.m.Range(func(k, v interface{}) bool {
		return f(k.(K), v.(V))
	})
}

// Len 获取当前存储的键值对数量
func (c *ConcurrentMap[K, V]) Len() int {
	// 返回计数器的值，并将 int64 转换为 int
	return int(atomic.LoadInt64(&c.length))
}

// Clear 清理所有键值对
func (c *ConcurrentMap[K, V]) Clear() {
	c.m.Range(func(key, value any) bool {
		c.m.Delete(key)
		return true
	})
	atomic.StoreInt64(&c.length, 0)
}

// Sizeof 当前并发数组内存大小
func (c *ConcurrentMap[K, V]) Sizeof() int64 {
	size := unsafe.Sizeof(c.m)
	return int64(size)
}

func test_ConcurrentMap() {
	// 创建一个 ConcurrentMap，key 是 string，value 是 int
	cm := NewConcurrentMap[string, int]()

	// 设置 key-value 对
	cm.Set("one", 1)
	cm.Set("two", 2)
	/*
		// 获取值
		if value, ok := cm.Get("one"); ok {
			fmt.Printf("key: one, value: %d\n", value)
		} else {
			fmt.Println("key: one not found")
		}
	*/

	// 使用 LoadOrStore
	value, loaded := cm.LoadOrStore("three", 3)
	if loaded {
		fmt.Printf("key: three already loaded with value: %d\n", value)
	} else {
		fmt.Printf("key: three stored with value: %d\n", value)
	}

	// 遍历所有的 key-value 对
	cm.Range(func(key string, value int) bool {
		fmt.Printf("key: %s, value: %d\n", key, value)
		return true
	})

	// 删除一个 key
	cm.Delete("one")
	/*
		// 确认删除
		if _, ok := cm.Get("one"); !ok {
			fmt.Println("key: one successfully deleted")
		}
	*/
}

/*
// ConcurrentMap 是一个并发安全的泛型 map
type ConcurrentMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

// NewConcurrentMap 创建一个新的 ConcurrentMap
func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		m: make(map[K]V),
	}
}

// Set 设置 key-value 对
func (c *ConcurrentMap[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value

}

// Get 获取 key 对应的 value
func (c *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	return value, ok
}

// Delete 删除 key 对应的 value
func (c *ConcurrentMap[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

// Has 检查 key 是否存在
func (c *ConcurrentMap[K, V]) Has(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.m[key]
	return ok
}

// Len 返回 map 的长度
func (c *ConcurrentMap[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.m)
}

// Iterate 遍历 map
func (c *ConcurrentMap[K, V]) Iterate(f func(K, V)) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for key, value := range c.m {
		f(key, value)
	}
}*/
