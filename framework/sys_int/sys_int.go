package sys_int

import (
	"strconv"
	"strings"
)

/*
	类型转换
	*
	* added by yh @ 2023-11-4 18:30

*注意:
*/

// 将 int32 切片转换为字符串
func Int32SliceToString(slice []int32) string {
	if slice == nil {
		return ""
	}
	// 将 int32 切片转换为字符串切片
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = strconv.Itoa(int(v))
	}

	// 将字符串切片连接为一个字符串
	result := ""
	for _, str := range strSlice {
		result += str
	}

	return result
}

/*判断字符串是否为nil 空字符串*/
func StringIsNullOrEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

/*字符串转Int32*/
func StringToInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

/*字符串转Int32*/
func StringToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}
