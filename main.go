package main

import (
	"gin_server_v1/helper"
	"gin_server_v1/middleware"
	_ "gin_server_v1/routes"
	_ "gin_server_v1/lib"
)


func main() {
	middleware.Server.Run(helper.Config.GetString("DingTalk.ServerIp")+":"+helper.Config.GetString("DingTalk.ServerPort"))
}