package sys_net

import "github.com/golang/protobuf/proto"

/*
请求处理程序 基类
added by yh @ 2023-04-23
*/

type IRequestHandler interface {
	//# 打印请求日志
	//PrintRequest()
	//# 验证数据来源以及当前状态的正确性
	//VerifyData() bool
	//处理请求
	Process()
	//# 储存数据
	//SaveData()
	//# 打印日志
	//PrintLog()
	//#打印验证失败日志
	//PrintVerifyFailLog()
	//# 打印返回值
	//PrintResponse()
	//# 返回结果
	GetResponse() *Resp
}

/*
	func (t *xxxx) GetResponse() *sys_net.Resp {
		//TODO implement me 这里cp
		return &sys_net.Resp{
			RetId: t.RequestId,
			Ret:   t.ret,
		}
	}
*/
type Resp struct {
	RetId int32
	Ret   proto.Message
}
type RequestHandler struct {
	RequestId int32
	Error     error
}

// # 打印请求日志
func (t *RequestHandler) PrintRequest() {
}

// 数据验证 默认过
func (t *RequestHandler) VerifyData() bool {
	return true
}

// # 储存数据
func (t *RequestHandler) SaveData() {

}

// # 打印日志
func (t *RequestHandler) PrintLog() {

}

// #打印验证失败日志
func (t *RequestHandler) PrintVerifyFailLog() {

}

// # 打印返回值
func (t *RequestHandler) PrintResponse() {

}

/*
// # 返回结果
func (t *RequestHandler) GetResponse() (string, proto.Message, error) {

	return "", nil, nil
}
*/
