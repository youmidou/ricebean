package sys_engine

type ERMIDispatchStatus int32

const (
	DispatchOK                   ERMIDispatchStatus = iota //成功
	DispatchObjectNotExist                                 //对象不存在
	DispatchMethodNotExist                                 //方法不存在
	DispatchParamError                                     //参数错误
	DispatchOutParamBuffTooShort                           //输出参数缓冲区太小
	SessionDisconnect                                      //会话断开
)
const MAX_IDENTITY_LEN = 32

/*
	RMI对象

added by yh @ 2024/05/11 11:53
*/
type RMIObject struct {
}

func NewRMIObject() *RMIObject {
	t := &RMIObject{}
	return t
}

/*
	RMI代理对象

added by yh @ 2024/05/11 11:53
*/
type RMIProxyObject struct {
	m_session Session
}

func NewRMIProxyObject() *RMIProxyObject {
	t := &RMIProxyObject{}
	return t
}
func (t *RMIProxyObject) __bind_session(session *Session) {

}
