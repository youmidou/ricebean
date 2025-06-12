package sys_random

import (
	"math/rand"
	"time"
)

// 生成在 [min, max] 范围内的随机整数
func RangeIntRandom(min int, max int) int {
	//min := 10
	//max := 50
	randomInRange := min + rand.Intn(max-min+1) // 生成在 [min, max] 范围内的随机整数
	return randomInRange
}

// 生成一个随机整数
func IntRandom(max int) int {
	randomInt := rand.Intn(max) // 生成 0 到 99 之间的随机整数
	return randomInt
}

func RandomValues(maxV, buildNum int32) []int32 {
	var values []int32
	for i := 0; i < int(buildNum); i++ {
		values = append(values, generateRandomNumber(maxV))
	}
	return values
}

func generateRandomNumber(maxV int32) int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31n(maxV)
}
