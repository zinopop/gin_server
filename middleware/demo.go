package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Demo struct {
}

func (*Demo) Test() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("我是demo中间件Test")
	}
}
