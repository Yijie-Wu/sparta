package cmd

import (
	"fmt"
	"sparta/conf"
	"sparta/database"
	"sparta/global"
	"sparta/logger"
	"sparta/router"
	"sparta/utils"
)

func StartApplication() {
	var initErr error
	conf.InitConfig()
	global.Logger = logger.InitLogger()

	db, err := database.InitDB()
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	global.DB = db

	redisClient, err := database.InitRedis()
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	global.RedisClient = redisClient

	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(any(fmt.Sprintf("Init database failed at:%s", initErr.Error())))
	}

	router.InitRouters()

	global.Logger.Info("Start app success")
}

func StopApplication() {
	global.Logger.Info("Stop app success")
}
