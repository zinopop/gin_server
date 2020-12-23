package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func init() {
	Server = gin.Default()
	fmt.Println("路由中间件初始化成功.....")
}
