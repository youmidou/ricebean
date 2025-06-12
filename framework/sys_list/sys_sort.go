package sys_list

import "sort"

type MyObject struct {
	Name  string
	Value int
}

func xxxx() {
	// 一组包含对象的数组
	myObjects := []MyObject{
		{Name: "Object1", Value: 5},
		{Name: "Object2", Value: 2},
		{Name: "Object3", Value: 8},
		{Name: "Object4", Value: 1},
		{Name: "Object5", Value: 7},
		{Name: "Object6", Value: 3},
		{Name: "Object7", Value: 9},
		{Name: "Object8", Value: 4},
		{Name: "Object9", Value: 6},
	}

	// 使用 sort 包中的 Sort 函数进行排序
	sort.Slice(myObjects, func(i, j int) bool {
		return myObjects[i].Value > myObjects[j].Value
	})
}
