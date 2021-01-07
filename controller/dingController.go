package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DingController struct {
	BaseController
}

func (controller DingController) DingSend(c *gin.Context) {
	params := struct {
		Text string `form:"text" json:"text" xml:"text"  binding:"required"`
	}{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "测试错误44444444"})
		return
	}
	//msg, err := helper.SendDingMsg(string(helper.ContextPush(params.Text)))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//controller.c.ShouldBind()
	c.JSON(200, gin.H{
		"data":    controller,
		"message": params.Text,
	})
}
