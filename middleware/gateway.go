package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Gateway struct {
}

func (*Gateway) Test() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("我是中间件Test")
		context.JSON(500, "aaasd")
		//context.Abort() //阻止继续执行
		context.Next()
	}
}
