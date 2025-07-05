package lobby_svc

import (
	pitaya "ricebean/pkg"
	"ricebean/pkg/modules"
)

func NewActivityModule(app pitaya.Pitaya) *ActivityModule {
	return &ActivityModule{app: app}
}

type ActivityModule struct {
	modules.Base
	app pitaya.Pitaya
}
