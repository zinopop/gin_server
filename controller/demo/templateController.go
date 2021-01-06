package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TemplateController struct {
}

func (TemplateController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Posts",
	})
}
