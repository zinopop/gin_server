package middleware

import (
	"fmt"
	"gin_server/helper"
	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func init() {
	gin.SetMode(helper.Config.GetString("Server.Mode"))
	Server = gin.Default()
	//debug release test
	fmt.Println("路由中间件初始化成功.....")
}
