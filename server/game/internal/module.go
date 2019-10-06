package internal

import (
	"github.com/irrelevantdotcom/leaf/module"
	"github.com/irrelevantdotcom/leafserver/server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
