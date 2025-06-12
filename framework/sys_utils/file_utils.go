package sys_utils

import (
	"fmt"
	"os"
	"phoenix-tudou/framework/log4"
)

// 将数据写入文件
// "output.txt
func WriteFile(filename string, data []byte) bool {
	/*
		// 创建文件
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer file.Close()
	*/
	// 写入数据
	err1 := os.WriteFile(filename, data, 0644)
	if err1 != nil {
		fmt.Println("写入文件失败："+filename, err1)
		return false
	}
	return true
}
func ReadFile(filename string) ([]byte, bool) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log4.Error(fmt.Sprintf("读取文件失败：", err))
		return nil, false
	}
	return data, true
}
