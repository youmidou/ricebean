package main

import (
	"flag"
	"fmt"
	"ricebean/_ymd"
	pitaya "ricebean/pkg"
	"ricebean/pkg/acceptor"
	"ricebean/pkg/config"
	"ricebean/pkg/conn/codec"
	"ricebean/pkg/conn/message"
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
	//数据包->解码器
	builder.PacketDecoder = _ymd.NewPomeloPacketDecoder()
	//数据包->编码器
	builder.PacketEncoder = codec.NewPomeloPacketEncoder()
	//消息->编码器
	builder.MessageEncoder = message.NewMessagesEncoder(cfg.Handler.Messages.Compression)
	builder.Serializer = protobuf.NewSerializer() // 设置 Protobuf 序列化器

	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", 1250))
	builder.AddAcceptor(tcp)
	//----------------------------------------------
	
	app = builder.Build()
	defer app.Shutdown()

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
