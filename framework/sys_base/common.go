package sys_base

import (
	"encoding/binary"
	"fmt"
	"phoenix-tudou/z_Tools/ProtoToCS/Protocal/pb"
	"reflect"
	"sync"
	"time"
)

var m_Locker *sync.RWMutex = new(sync.RWMutex)

func Lock() {
	m_Locker.Lock()
}
func Unlock() {
	m_Locker.Unlock()
}

// 获取消息模块
func GetCmdModule(msgId int32) pb.CmdModule {
	quotient := msgId / 1000
	return pb.CmdModule(quotient)
}

// 协议转换 协议id除1000 取整 模块id ,取余数 消息id
func ProtocolIdAnalysis(msgId int32) {
	//
	//msgId = 1003002
	//1001001
	//1001002
	// 整数相除并取整
	quotient := msgId / 1000
	fmt.Println("整数相除并取整:", quotient)

	// 取余数
	remainder := msgId % 1000
	fmt.Println("取余数:", remainder)
}

/*
type RequestInfo struct {}
usage:
requestInfo := createInstance("RequestInfo").(RequestInfo)
*/
func CreateInstance(typeName string) interface{} {
	// get the type by name
	t := reflect.TypeOf(typeName)
	// create a new instance of the type
	instance := reflect.New(t).Elem().Interface()
	return instance
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// 字节转换成整形 binary.BigEndian 反序 binary.LittleEndian
func BytesToInt(data []byte) int {
	buff := make([]byte, 4)
	copy(buff, data)
	tmp := int32(binary.BigEndian.Uint32(buff))
	return int(tmp)
}

// 整形16转换成字节
func Int16ToBytes(val int16) []byte {
	tmp := uint16(val)
	buff := make([]byte, 2)
	binary.BigEndian.PutUint16(buff, tmp)
	return buff
}

func GetUShort(data []byte) uint16 {
	buff := make([]byte, 2)
	copy(buff, data)
	var tmp = uint16(binary.BigEndian.Uint16(buff))
	return uint16(tmp)
}

func GetDBTime(strTime string) *time.Time {
	DefaultTimeLoc := time.Local
	loginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, DefaultTimeLoc)
	return &loginTime
}
