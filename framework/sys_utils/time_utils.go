package sys_utils

import (
	"fmt"
	"time"
)

// 判断时间是否属于今天
func IsToday(t time.Time) bool {
	if t.IsZero() {
		return false
	}
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// 返回 -1 表示 t1 比 t2 小，1 表示 t1 比 t2 大，0 表示两个时间相等
func Time_CompareTime(t1, t2 time.Time) int {
	if t1.Before(t2) {
		return -1 //t1<t2
	} else if t1.After(t2) {
		return 1 //t1>t2
	} else {
		return 0
	}
}

// 时间转换工具
func TimeToInt64(t time.Time) int64 {
	//t := time.Now().Unix()
	return t.Unix()
}

//区域时间转换
//Go语言可以将时间转换为地区时间。您可以使用time包中的函数进行转换，例如time.LoadLocation()和time.In()。
//time.Now().Format("YYYY-MM-DD HH:mm:ss.sss")
//time.Now().Format("2006-01-02 15:04:05.000")
//减少
//   t = t.Add(-10 * time.Second)
//添加
//t3 := t1.Add(time.Hour * 24)
//比较 t2.Sub(t3)

// 比较两个时间大小 now < later
func Before(now time.Time, later time.Time) {
	if now.Before(later) {
		fmt.Println("now is before later")
	}
}

// now > later
func After(now time.Time, later time.Time) {
	if later.After(now) {
		fmt.Println("later is after now")
	}
}
func Equal(now time.Time, later time.Time) {
	if !now.Equal(later) {
		fmt.Println("now and later are not equal")
	}
}
