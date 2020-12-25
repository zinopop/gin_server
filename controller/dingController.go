package controller

import (
	"fmt"
	"gin_server/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DingController struct {
}

type SendMsg struct {
	Text string `form:"text" json:"text" xml:"text"  binding:"required"`
}

//
func (DingController) DingSend(c *gin.Context) {
	var SendMsg SendMsg
	if err := c.ShouldBind(&SendMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "测试错误"})
		return
	}
	msg, err := helper.SendDingMsg(string(helper.ContextPush(SendMsg.Text)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data":    "",
		"message": msg,
	})
}

//test
func (DingController DingController) Test(a string) string {
	fmt.Println(a)
	return a
}
