package sys_list

import "testing"

func TestStore(s *testing.T) {

	l := NewList()
	l.Foreach(func(k any, v any) {

	})
	l.For(func(k any, v any) {

	})

	l.Add(1)
	l.Add(2)
	l.Add(33)
	l.Add(4)
	l.Add(5)
	l.Remove(4)
	//l.RemoveIndex()
	//l.GetItemByIndex()
	/*
		l.SortSlice(func(a, b int) bool {
			return l[a] > l[b]
		})
	*/
}
