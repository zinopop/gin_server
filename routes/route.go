package routes

import (
	"gin_server/controller"
	andoncloud_v6_controller "gin_server/controller/andoncloud/v6"
	"gin_server/middleware"
)

func init() {
	//钉钉消息推送(普通消息)
	middleware.Server.Any("/sendMsg", controller.DingController{}.DingSend)
	//webhook自动构建
	middleware.Server.GET("/webhook/rebuild", controller.WebhookController{}.Rebuild)
	//默认路由
	middleware.Server.Any("/default", controller.DefaultController{}.Index)

	//九安医疗路由分组
	andoncloud := middleware.Server.Group("andoncloud")
	andoncloud_v6 := andoncloud.Group("v6")
	{
		//测试路由
		andoncloud_v6.POST("/test", andoncloud_v6_controller.DingController{}.Test)
	}
}
