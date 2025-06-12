package sys_base

import "fmt"

/*
	添加宏定义

added by yh @ 2024/04/23 10:24
列子:
*/
func assert(isOk bool, msg string) {
	if isOk == false {
		str := fmt.Sprintf("assert failed %s", msg)
		panic(str)
	}
}
