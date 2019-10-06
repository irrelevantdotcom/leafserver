package main

import (
	"github.com/irrelevantdotcom/leaf"
	lconf "github.com/irrelevantdotcom/leaf/conf"
	"github.com/irrelevantdotcom/leafserver/server/conf"
	"github.com/irrelevantdotcom/leafserver/server/game"
	"github.com/irrelevantdotcom/leafserver/server/gate"
	"github.com/irrelevantdotcom/leafserver/server/login"
)	

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
