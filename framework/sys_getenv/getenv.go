package sys_getenv

import (
	"fmt"
	"os"
	"runtime"
)

/*
获取系统环境
added by yh @ 2023-04-26
*/

// 是否是调试模式 true:调试模式; false:prod发布模式;
var Debug = true

const (
	MacOS   = "darwin"
	Linux   = "linux"
	windows = "windows"
)

func CurrentSystem() string {
	// 获取当前操作系统
	os := runtime.GOOS
	return os
}
func IsMacOS() bool {
	// 获取当前操作系统
	os := runtime.GOOS
	return os == "darwin"
}
func IsLinux() bool {
	// 获取当前操作系统
	os := runtime.GOOS
	return os == "linux"
}
func IsWindows() bool {
	// 获取当前操作系统
	os := runtime.GOOS
	return os == "windows"
}

func Init() {
	//isDebug()
	// 获取当前操作系统
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Running on macOS")
		Debug = true
	case "linux":
		Debug = false
		fmt.Println("Running on Linux")
	case "windows":
		Debug = true
		fmt.Println("Running on Windows")
	default:
		Debug = true
		fmt.Println("Running on unknown operating system:", os)
	}
}

// 默认debug prod
func isDebug() {
	str := os.Getenv("Debug")
	if str == "" {
		println("Debug mode on Debug = true")
		Debug = true
	} else {
		println("No Debug Debug = false")
		Debug = false
	}
}
