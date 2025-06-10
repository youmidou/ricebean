package _interfaces

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/topfreegames/pitaya/v3/pkg/component"
)

type LobbySvc interface {
	component.Component
	SendRPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message) error

	//SendToUserGameWorld(sceneId int32, userId _def.UserID, data []byte) bool
	//SendToGateWayAllUser(data []byte) bool                //
	//SendToSceneGameWorld(sceneId int32, data []byte) bool //发送场景消息

	GetUserManager() GlobalUserManager
	GetActivityManager() ActivityManager
	GetRankManager() RankManager
	//SendToUserGameWorld() bool
	//SendToGateWayAllUser() bool
	//SendToSceneGameWorld() bool
	//SendToAllGameWorld() bool
}
