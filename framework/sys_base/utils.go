/*
登录模块 共用函数
added by yh @ 2023/5/25 13:23
注意:
*/
package sys_base

import (
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/framework/sys_json"
)

/*
登陆凭证数据
规则 DeviceId:必须有 剩下的平台登陆必须有一个 代表什么平台登陆
*/
type CredentialInfo struct {
	UserId     string //服务器设备唯一标识
	DeviceId   string //客户端设备Id
	FacebookId string //绑定fb
	AppleSubId string
	HmsId      string
}

// ************运算逻辑****************************************
func SetCredentialInfo(info *CredentialInfo) string {
	// 将Go结构体转换为JSON数据
	//jsonBytes, err := sys_json.Marshal(info)
	jsonBytes, err := sys_json.MarshalToString(info)
	if err != nil {
		log4.Error("setDataByCredential->json %s", err.Error())
	}
	//var credential = sys_utils.Encrypt(jsonBytes)
	return jsonBytes
}
func GetCredentialInfo(credential string) *CredentialInfo {

	//credential = sys_utils.Decrypt(credential)
	info := &CredentialInfo{}
	if credential == "" {
		return info
	}
	err := sys_json.Unmarshal([]byte(credential), info)
	if err != nil {
		log4.Error("credential-> %s", err.Error())
	}
	return info
}
