package main

import (
	"gin_server/helper"
	"gin_server/lib"
	"gin_server/middleware"
	"gin_server/model"
	_ "gin_server/routes"
)

func main() {

	db := lib.GetConn()

	//同步表结构
	db.Sync(&model.UserTest{})

	middleware.Server.Run(helper.Config.GetString("Server.Http.ServerIp") + ":" + helper.Config.GetString("Server.Http.ServerPort"))
}
