package cmd

import (
	"sparta/conf"
	"sparta/global"
	"sparta/router"
)

func StartApplication() {
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	router.InitRouters()
	global.Logger.Info("Start app success")
}

func StopApplication() {
	global.Logger.Info("Stop app success")
}
