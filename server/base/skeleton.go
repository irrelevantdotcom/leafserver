package base

import (
	"github.com/irrelevantdotcom/leaf/chanrpc"
	"github.com/irrelevantdotcom/leaf/module"
	"github.com/irrelevantdotcom/leafserver/server/conf"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
