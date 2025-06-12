/*
	封装数组

added by yh @ 2023/05/04 18:23
注意:
*/
package sys_list

import "sort"

type List []any

func NewList() List {
	t := List{}
	return t
}

/*
//list := list.New()
//list.PushBack(cid)
*/
func (t *List) Add(item any) {
	*t = append(*t, item)
}

func (t *List) Remove(item any) bool {
	if item == nil {
		return false
	}
	for i, a := range *t {
		if a == item {
			t.RemoveIndex(i)
			return true
		}
	}
	return false
}
func (t *List) RemoveIndex(index int) {
	*t = append((*t)[:index], (*t)[index+1:]...)
}

func (t *List) Len() int {
	return len(*t)
}

// 切片排序
func (t *List) SortSlice(less func(a, b int) bool) {
	/*
		// 使用 sort 包中的 Sort 函数进行排序
		sort.Slice(t, func(i, j int) bool {
			return t[i].Value > myObjects[j].Value
		})
	*/
	sort.Slice(t, less)
}

func (t *List) Foreach(f func(k any, v any)) {
	for i, a := range *t {
		f(i, a)
	}
}

func (t *List) For(f func(k any, v any)) {
	for i, a := range *t {
		f(i, a)
	}
}
