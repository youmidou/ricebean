package sys_net

import (
	"github.com/golang/protobuf/proto"
	"phoenix-tudou/framework/log4"
)

// 包-解析
func PacketDecoder() {

}

// 包-编码
func PacketEncoder() {

}

// 消息-解析
func MessageDecoder() {

}

// 消息-编码
func MessageEncoder() {

}

// 封装包 消息包
func MessagePack(msgId int32, data []byte) []byte {
	b := NewBytebuffer()
	b.WriteInt32(msgId)
	b.WriteBytes(data)
	return b.ToBytes()
}

// 解压消息包
func MessageUnpack(data []byte) (int32, []byte) {
	buff := NewBytebuffer()
	buff.SetData(data)
	var msgId = buff.ReadInt32() //buff.ReadString()
	var bytes = buff.ReadBytes()
	return msgId, bytes
}

// 封装消息内容
func MessageContentPack(msg proto.Message) ([]byte, error) {
	// 将对象转换为二进制数据
	data, err := proto.Marshal(msg)
	if err != nil {
		log4.Error("marshaling protobuf=>%s", err.Error())
	}
	return data, err
}

/*
解码列子:
var msg =&pb.Net_xxxx{}
sys_net.MessageContentUnpack(buff,msg)
*/
func MessageContentUnpack(msgId int32, buff []byte, msg proto.Message) error {
	//将二进制解析成对象
	err := proto.Unmarshal(buff, msg)
	if err != nil {
		log4.Error("MessageContentUnpack->msgId=%d err=%d", msgId, err.Error())
	}
	return err
}

/*
// 检查作弊
func CheckCtxCheat(ctx *CTXConn, reqName string) {

}
*/
