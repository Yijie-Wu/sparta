package cmd

import (
	"fmt"
	"sparta/conf"
	"sparta/router"
)

func StartApplication() {
	fmt.Println("Start app")
	conf.InitConfig()
	router.InitRouters()
}

func StopApplication() {
	fmt.Println("Stop app")
}
