/*
*
added by yh @ 2023/5/13 10:10
注意:
*/
package sys_utils

import "phoenix-tudou/framework/log4"

// 宏定义
func Assert(isbool bool, message string, a ...any) {
	if isbool == false {
		log4.Error(message, a)
	}
}

// ContainsInt32 判断一个数组内是否包含一个int32元素
func ContainsInt32(slice []int32, item int32) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
