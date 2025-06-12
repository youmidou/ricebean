package sys_bit

/*
	位运算

added by yh @ 2023/12/05 15:12
注意:
*/
//--------todo  int32---------------------------------------------------------
func GetInt32BitList(target int32) []int32 {
	// 创建一个位数组切片
	bits := make([]int32, 32)
	// 使用位运算获取每一位的值
	for i := 0; i < 32; i++ {
		bits[i] = int32((target >> i) & 1)
	}

	// 返回位数组
	return bits
}
func GetInt32Bit(bits []int32) int32 {
	var result int32
	for i, bit := range bits {
		result |= bit << i
	}
	return result
}

// 将一个指定位置的位设置为指定的值 bitValue:0或者1
func SetInt32BitValue(target int32, pos int32, bitValue int8) int32 {
	// 清除指定位置的位
	target = target & ^(1 << pos)
	// 设置指定位置的位为指定值
	if bitValue == 1 {
		target = target | (1 << pos)
	} else {
		target = target | (0 << pos)
	}
	return target
}

// pos:0-31
func GetInt32BitValue(target int32, pos int32) (int32, bool) {
	if 0 <= pos && pos < 32 {
		// 取第4位的值
		fourthBit := (target >> pos) & 1
		return fourthBit, true
	}
	return 0, false
}

//--------todo int64---------------------------------------------------------

// 将 int64 转换为位数组
func GetInt64BitList(target int64) []int32 {
	// 创建一个位数组切片
	bits := make([]int32, 64)
	// 使用位运算获取每一位的值
	for i := 0; i < 64; i++ {
		bits[i] = int32((target >> i) & 1)
	}
	// 返回位数组
	return bits
}

// 将位数组转换为 int64 数字
func GetInt64Bit(bits []int32) int64 {
	var result int64

	for i, bit := range bits {
		result |= int64(bit) << i
	}

	return result
}

// pos:0-63 返回0或者1
func GetInt64BitValue(target int64, pos int32) (int32, bool) {
	if 0 <= pos && pos < 63 {
		// 取第4位的值
		fourthBit := (target >> pos) & 1
		return int32(fourthBit), true
	}
	return 0, false

}
func SetInt64BitValue(target int64, pos int32, bitValue int8) int64 {
	// 清除指定位置的位
	target = target & ^(1 << pos)
	// 设置指定位置的位为指定值
	if bitValue == 1 {
		target = target | (1 << pos)
	} else {
		target = target | (0 << pos)
	}
	return target
}

// --------todo xxx---------------------------------------------------------
// 从一个整数中读取指定位置的位 返回位值 0或1
func ReadBit(target int32, pos int32) int32 {
	bit := (target >> pos) & 1
	return bit
}
