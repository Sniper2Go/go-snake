package mgr

import (
	"go-snake/accountServer"
	"go-snake/app/application"
	"go-snake/gameServer"
	"go-snake/gateServer"
	"go-snake/loginServer"
	"go-snake/robot"
)

var (
	apps = map[string]application.ApplicationIF{
		CstGame:    gameServer.New(),
		CstGate:    gateServer.New(),
		CstLogin:   loginServer.New(),
		CstAccount: accountServer.New(),
		CstRobot:   robot.New(),
	}
)

func GetApp(sn string) application.ApplicationIF {
	return apps[sn]
}
