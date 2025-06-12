package sys_runtime

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"runtime"
)

// 获取可用内存
func GetMemory() string {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	//---可用内存将字节转换为兆字节------------------------------------
	availableMemoryMB := float64(memStats.Sys) / 1024 / 1024
	//fmt.Printf("可用内存: %.2f MB\n", availableMemoryMB)

	//----当前进程占用的总内存大小--------------------------------------
	totalMemory := memStats.TotalAlloc
	// 将字节转换为兆字节
	totalMemoryMB := float64(totalMemory) / 1024 / 1024
	//fmt.Printf("当前进程占用的总内存: %.2f MB\n", totalMemoryMB)
	//--------------------------------------------------------------

	str := fmt.Sprintf("可用内存: %.2f MB\n 当前进程占用的总内存: %.2f MB\n",
		availableMemoryMB, totalMemoryMB)
	//memStats.BuckHashSys
	return str
}
func GetMemInfo() string {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("get memory info fail. err： ", err)
	}
	// 获取总内存大小，单位GB
	memTotal := memInfo.Total / 1024 / 1024 / 1024
	// 获取已用内存大小，单位MB
	memUsed := memInfo.Used / 1024 / 1024
	// 可用内存大小
	memAva := memInfo.Available / 1024 / 1024
	// 内存可用率
	memUsedPercent := memInfo.UsedPercent

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	usedMemory := (m.Sys - m.HeapReleased) / 1024 / 1024

	str := fmt.Sprintf("总: %v GB,应用已用%v MB 总用: %v MB, 可用: %v MB, 使用率: %.3f",
		memTotal, usedMemory, memUsed, memAva, memUsedPercent)

	return str
}

func xxxx() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Sys 字段表示已分配的内存总量，包括当前使用中的内存和被释放但未返回给操作系统的内存。
	// 此处我们获取已分配的内存总量，然后减去未返回给操作系统的内存量，得到当前进程使用的内存量。
	usedMemory := m.Sys - m.HeapReleased

	// 打印当前进程使用的内存量
	fmt.Printf("当前进程使用的内存: %v bytes\n", usedMemory)
}
