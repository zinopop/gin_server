package routes

import (
	"gin_server_v1/controller"
	"gin_server_v1/middleware"
)

func init() {
	//钉钉消息推送(普通消息)
	middleware.Server.Any("/sendMsg", controller.DingSend)
	//webhook自动构建
	middleware.Server.GET("/webhook/rebuild", controller.Rebuild)
}
