package controller

import (
	"gin_server_v1/lib"
	"github.com/gin-gonic/gin"
)

func Rebuild(c *gin.Context) {
	lib.RebuildTask()
}
