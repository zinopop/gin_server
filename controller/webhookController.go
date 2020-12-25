package controller

import (
	"gin_server/lib"
	"github.com/gin-gonic/gin"
)

type WebhookController struct {
}

func (WebhookController) Rebuild(c *gin.Context) {
	lib.RebuildTask()
}
