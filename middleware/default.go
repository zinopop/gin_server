package middleware

import "github.com/gin-gonic/gin"

var Server *gin.Engine

func init()  {
	Server = gin.Default()
}