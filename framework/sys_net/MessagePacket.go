package sys_net

import (
	"encoding/binary"
	"fmt"
)

// 消息包接收
func SendMessagePacket(msgId int32, msgData []byte) []byte {
	//将 msgId 转换成[]byte
	msgIdBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgIdBytes, uint32(msgId))
	// 将 msgIdBytes 和 content 组合成一个新的 []byte 切片 result
	content := append(msgIdBytes, msgData...)
	var l int32 = int32(len(content))
	l_btye := make([]byte, 4)
	binary.LittleEndian.PutUint32(l_btye, uint32(l))
	//--封包---
	msgPkg := append(l_btye, content...)
	return msgPkg
}

// 消息包接收
func ReceiveMessagePacket(data []byte) (int32, []byte, error) {
	// 如果接收到的数据不够，返回错误
	if len(data) < 8 {
		return 0, nil, fmt.Errorf("insufficient data length")
	}

	// 解析长度字段
	length := binary.LittleEndian.Uint32(data[:4])
	// 如果接收到的数据长度不够，继续等待更多数据到达
	if len(data) < int(length) {
		return 0, nil, fmt.Errorf("insufficient data length")
	}
	// 解析消息号字段
	msgId := int32(binary.LittleEndian.Uint32(data[4:8]))
	// 解析内容
	content := data[8:]
	return msgId, content, nil
}
