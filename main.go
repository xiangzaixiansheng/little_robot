package main

import (
	"little_robot/conf"
	"little_robot/routes"
)

func main() {
	// Ek1+Ep1==Ek2+Ep2
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
