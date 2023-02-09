package main

import (
	"little_robot/conf"
	"little_robot/global"
	util "little_robot/pkg/utils"
	"little_robot/robot"
	"little_robot/routes"
	"time"
)

func main() {
	begin := time.Now()
	defer func() {
		util.LogrusObj.Infoln("[INFO]本次已经运行: %s", time.Since(begin).String())
	}()
	conf.Init()
	r := routes.NewRouter()
	//初始化map
	global.InitGlobalWechatBotMap()
	util.LogrusObj.Infoln("走到这里了")
	//首次登陆
	robot.First_load()

	_ = r.Run(conf.HttpPort)
}
