package sys_utils

//两个二进制数组 合并成一个数组
func Append(a[]byte,b[]byte) []byte {

	return append(a,b...)
}