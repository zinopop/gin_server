package main

import (
	"gin_server/helper"
	"gin_server/middleware"
	_ "gin_server/routes"
	"log"
)

func main() {
	//db := lib.GetConn()
	//
	////同步表结构
	//db.Sync(&model.UserTest{})
	//注册html
	middleware.Server.LoadHTMLGlob("assets/*")
	//注册前端静态资源
	middleware.Server.Static("/assets", "./assets")
	if err := middleware.Server.Run(helper.Config.GetString("Server.Http.ServerIp") + ":" + helper.Config.GetString("Server.Http.ServerPort")); err != nil {
		log.Fatal(err.Error())
	}
}
