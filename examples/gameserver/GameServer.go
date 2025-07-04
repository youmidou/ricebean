package main

import (
	"context"
	"flag"
	"fmt"
	"ricebean/_ymd_packet"
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

	//----------------------------------------------
	builder.Serializer = protobuf.NewSerializer() // 设置 Protobuf 序列化器
	//数据包编解码器
	builder.PacketCodec = codec.NewPomeloMessagePacket()
	//消息->编码器
	//builder.MessageCodec = message.NewYmdMessageCodec(cfg.Handler.Messages.Compression)
	builder.MessageCodec = _ymd_packet.NewYmdPacketEncoder(cfg.Handler.Messages.Compression)

	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", 1250))
	builder.AddAcceptor(tcp)
	//ws := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", 1251))
	//builder.AddAcceptor(ws)
	//注册网关接收模块
	//builder.SetGatewayHandlerSvc()
	//----------------------------------------------
	//GatewayHandlerService. gateway_svc

	app = builder.Build()
	defer app.Shutdown()

	//-----------大厅-----------------------
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

	//接收路由地址 server.service.handler
	err := app.SetDictionary(map[string]uint16{
		//---------Gateway------------------------------------------------
		"LoginSvc.OnGatewayReceive": 1, //
		"LobbySvc.OnGatewayReceive": 2, //
		"ThemeSvc.OnGatewayReceive": 3, //
		//"game.game.entergame":                  7,
	})

	if err != nil {
		fmt.Printf("error setting route dictionary %s\n", err.Error())
	}
	app.Start()

}
