package sys_math

import (
	"math/rand"
	"time"
)

var PerW int32 = 10000

// 获取一个随机值，默认是随机0到10000，最大值可以传递
func GetRandom(maxV ...int32) int32 {
	var maxValue int32
	if len(maxV) == 0 {
		maxValue = PerW
	} else {
		maxValue = maxV[0]
	}
	fValue := rand.Int31n(maxValue)
	return fValue
}

// 随机指定数量的随机数，默认最大值10000
func GetRandoms(num int32, maxV ...int32) []int32 {
	var maxValue int32
	if len(maxV) == 0 {
		maxValue = PerW
	} else {
		maxValue = maxV[0]
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	vs := make([]int32, num)
	var i int32 = 0
	for ; i < num; i++ {
		vs[i] = r.Int31n(maxValue)
	}
	return vs
}
