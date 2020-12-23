package main

import (
	"gin_server_v1/helper"
	_ "gin_server_v1/lib"
	"gin_server_v1/middleware"
	_ "gin_server_v1/routes"
)

func main() {
	middleware.Server.Run(helper.Config.GetString("Server.Http.ServerIp") + ":" + helper.Config.GetString("Server.Http.ServerPort"))
}
