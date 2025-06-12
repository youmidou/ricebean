package sys_string

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
func Split(str string, sep string) []string {
	// 使用空格作为分隔符分割字符串
	result := strings.Split(str, sep)
	return result
}

// 一维数组
// 二维数组 10007,3;10008,3;10009,3;10010,3;10011,3;10012,3
func TwoDimensionalArrayInt32(str string) [][]int32 {
	// 使用空格作为分隔符分割字符串
	result2 := strings.Split(str, ";")
	// 创建一个整数数组
	var twolist [][]int32
	for _, s2 := range result2 {
		if s2 == "" {
			continue
		}
		var onelist []int32
		result1 := strings.Split(s2, ",")
		for _, s := range result1 {
			if s == "" {
				continue
			}
			// 将字符串转换为整数
			num, _ := strconv.ParseInt(s, 10, 32)
			// 将整数添加到整数数组
			onelist = append(onelist, int32(num))
		}
		twolist = append(twolist, onelist)
	}
	return twolist
}

// 数组随机洗牌
func Shuffle(alist []any) {
	// 对切片进行洗牌
	rand.Shuffle(len(alist), func(i, j int) {
		alist[i], alist[j] = alist[j], alist[i]
	})
}

// "300,500,1000,2000"
func SplitToInt32(str string, sep string) []int32 {
	// 使用空格作为分隔符分割字符串
	result := strings.Split(str, sep)
	// 创建一个整数数组
	var intArray []int32
	// 使用循环将字符串转换为整数并添加到整数数组中
	for _, s := range result {
		if s == "" {
			continue
		}
		// 将字符串转换为整数
		num, _ := strconv.ParseInt(s, 10, 32)
		// 将整数添加到整数数组
		intArray = append(intArray, int32(num))
	}

	return intArray
}

func StringToInt(strNum string) (int, error) {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0, err
	}
	return num, err

}
