package proxy_svc

import (
	"ricebean/framework/sys_net"
	pitaya "ricebean/pkg"
	"ricebean/pkg/component"
)

func NewProxySvc(app pitaya.Pitaya) *ProxySvc {
	return &ProxySvc{app: app}
}

type ProxySvc struct {
	component.Base
	app          pitaya.Pitaya
	_request_map map[int32]func(msgId int32, bytes []byte) (sys_net.IRequestHandler, error)
}

func (t *ProxySvc) Init() {

}
