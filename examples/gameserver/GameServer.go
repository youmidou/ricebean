package main

import (
	"context"
	"flag"
	"fmt"
	"ricebean/examples/GameServer/lobby_svc"
	pitaya "ricebean/pkg"
	"ricebean/pkg/acceptor"
	"ricebean/pkg/agent"
	"ricebean/pkg/component"
	"ricebean/pkg/config"
	"ricebean/pkg/conn/codec"
	"ricebean/pkg/conn/message"
	"ricebean/pkg/route"
	"ricebean/pkg/serialize/protobuf"
)

var app pitaya.Pitaya

func main() {

	svType := flag.String("type", "Game", "the server type")
	serverMetadata := map[string]string{
		"ipaddress": "192.168.1.100:1250", // Replace with actual public IP
	}
	cfg := config.NewDefaultPitayaConfig()

	builder := pitaya.NewDefaultBuilder(true, *svType, pitaya.Cluster, serverMetadata, *cfg)
	//--------------初始化模块--------------------------------

	//--------------网关模块--------------------------------
	builder.Serializer = protobuf.NewSerializer() // 设置 Protobuf 序列化器
	//数据包编解码器
	builder.PacketCodec = codec.NewPomeloMessagePacket()
	//消息->编码器
	builder.MessageCodec = message.NewYmdMessageCodec(cfg.Handler.Messages.Compression)

	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", 1250))
	builder.AddAcceptor(tcp)

	//-----------初始化网关---------------------------------
	//注册网关接收模块
	//builder.SetGatewayHandlerSvc()
	//----------------------------------------------
	//GatewayHandlerService. gateway_svc

	app = builder.Build()
	defer app.Shutdown()
	//-----------功能模块--------------------------
	//lobby_svc.NewGlobalUserManager

	activityModule := lobby_svc.NewActivityModule(app)
	err := app.RegisterModule(activityModule, "activityModule")
	if err != nil {
		return
	}
	act, _ := app.GetModule("activityModule")
	cc := act.(*lobby_svc.ActivityModule)
	if cc != nil {

	}
	//-----------消息接收服务-----------------------
	lobbySvc := lobby_svc.NewLobbySvc(app)
	app.Register(lobbySvc, component.WithName("LobbySvc"))
	//app.RegisterRemote(lobbySvc, component.WithName("LobbySvc"))

	app.SetOnGatewayReceive(func(ctx context.Context, a agent.Agent, route *route.Route, msg *message.Message) {
		s := a.GetSession()
		s.OnClose(func() {
		})
		s.Bind(ctx, "")
		//s.Bind("1")
		lobbySvc.OnGatewayReceive(ctx, a, route, msg)
	})

	//接收路由地址 server.service.handler Game.LobbySvc.1001002 not found
	err = app.SetDictionary(map[string]uint16{
		//---------Gateway------------------------------------------------
		"LoginSvc.OnGatewayReceive": 1, //
		"LobbySvc.M1001002":         2, //
		//"ThemeSvc.OnGatewayReceive": 3, //
		//"game.game.entergame":                  7,
	})

	if err != nil {
		fmt.Printf("error setting route dictionary %s\n", err.Error())
	}
	app.Start()

}
