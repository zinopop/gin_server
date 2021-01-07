package routes

import (
	"gin_server/controller"
	controller_andoncloud_v6 "gin_server/controller/andoncloud/v6"
	controller_demo "gin_server/controller/demo"
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
	router_andoncloud_v6 := middleware.Server.Group("andoncloud").Group("v6")
	{
		//测试路由
		router_andoncloud_v6.POST("/test", controller_andoncloud_v6.DingController{}.Test)
	}

	//demo路由分组
	router_demo := middleware.Server.Group("demo")
	{
		//模板渲染
		router_demo.GET("template/index", controller_demo.TemplateController{}.Index)
		//接口返回
		router_demo.Any("api/getTest", controller_demo.ApiController{}.GetTest)
	}
}
