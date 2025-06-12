package sys_base

import (
	"github.com/bwmarrin/snowflake"
)

var _node *snowflake.Node

// 实例化
func SetUUIDInit(serverId int64) error {
	n, err := snowflake.NewNode(serverId)
	_node = n
	return err
}

// 获取
func GetUUID3() string {
	return _node.Generate().Base32()
}
func GetServerId(uuid string) int64 {
	id2, err := snowflake.ParseBase32([]byte(uuid))
	if err != nil {
		return -1
	}
	return id2.Node()
}
