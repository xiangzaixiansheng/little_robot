package main

import (
	"little_robot/conf"
	util "little_robot/pkg/utils"
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
	_ = r.Run(conf.HttpPort)
}
