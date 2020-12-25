package controller

import (
	"gin_server/lib"
	"github.com/gin-gonic/gin"
)

func Rebuild(c *gin.Context) {
	lib.RebuildTask()
}
