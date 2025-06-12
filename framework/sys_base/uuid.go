package sys_base

import "fmt"

type UUID string

const (
	ChanneRange = 10    //渠道范围
	ServerRange = 10000 //区范围
)

/*
	uuid

added by yh @ 2024/04/24 10:48
规则: channel string, pName string, serverId int32
userId:用户Id
channelId:渠道Id
serverId:区Id
*/
func GetUUID(userId int32, serverId int32, channelId int32) int64 {
	if 0 > channelId || channelId > ChanneRange {
		return -1
	}
	if 0 > serverId || serverId > ServerRange {
		return -2
	}
	//userId1001000
	channel := channelId             //0-99 //渠道
	server := serverId * ChanneRange //0-9999 //区
	userIdBit := int64(userId) * ServerRange * ChanneRange
	var uuid int64
	uuid = userIdBit + int64(server+channel)
	return uuid
}
func GetUUIDToString(userId int32, serverId int32, channelId int32) string {
	str := fmt.Sprintf("%d", GetUUID(userId, serverId, channelId))
	return str
}
func SetUUIDToUserIDServerIdChannelId(uuid int64) (int32, int32, int32) {
	userId := uuid / (ServerRange * ChanneRange)
	yu := uuid % (ServerRange * ChanneRange)
	serverId := yu / ChanneRange
	yu2 := yu % ChanneRange
	channel := yu2
	return int32(userId), int32(serverId), int32(channel)
}
