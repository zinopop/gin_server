package main

import (
	"gin_server_v1/helper"
	"gin_server_v1/lib"
	_ "gin_server_v1/lib"
	"gin_server_v1/middleware"
	"gin_server_v1/model"
	_ "gin_server_v1/routes"
)

func main() {

	db := lib.GetConn()

	//同步表结构
	db.Sync(&model.UserTest{})

	middleware.Server.Run(helper.Config.GetString("Server.Http.ServerIp") + ":" + helper.Config.GetString("Server.Http.ServerPort"))
}
