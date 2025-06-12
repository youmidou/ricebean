package sys_thread

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

// 获取当前携程Id
func GetGoroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := bytes.Fields(buf[:n])[1]
	id, err := strconv.ParseInt(string(idField), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
