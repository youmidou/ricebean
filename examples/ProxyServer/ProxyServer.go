package ProxyServer

import (
	"flag"
	"ricebean/examples/ProxyServer/proxy_svc"
	pitaya "ricebean/pkg"
	"ricebean/pkg/component"
	"ricebean/pkg/config"
	"ricebean/pkg/conn/codec"
	"ricebean/pkg/conn/message"
	"ricebean/pkg/serialize/protobuf"
)

var app pitaya.Pitaya

func main() {
	//负载均衡代理服务器

	svType := flag.String("type", "Proxy", "the server type")
	serverMetadata := map[string]string{
		"ipaddress": "192.168.1.100:1250", // Replace with actual public IP
	}
	cfg := config.NewDefaultPitayaConfig()

	builder := pitaya.NewDefaultBuilder(false, *svType, pitaya.Cluster, serverMetadata, *cfg)
	//数据包编解码器
	builder.PacketCodec = codec.NewPomeloMessagePacket()
	//消息->编码器
	builder.MessageCodec = message.NewPomeloPacketEncoder(cfg.Handler.Messages.Compression)
	builder.Serializer = protobuf.NewSerializer() // 设置 Protobuf 序列化器

	//builder.AddGin()
	app = builder.Build()
	defer app.Shutdown()
	//-----------proxy-----------------------
	lobbySvc := proxy_svc.NewProxySvc(app)
	app.Register(lobbySvc, component.WithName("ProxySvc"))

}
