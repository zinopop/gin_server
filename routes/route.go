package routes

import (
	"gin_server_v1/controller"
	"gin_server_v1/middleware"
)

func init()  {
	//钉钉消息推送(普通消息)
	middleware.Server.POST("/sendMsg", controller.DingSend)
	//测试数据库
	middleware.Server.POST("/dbTest", controller.DbText)
}