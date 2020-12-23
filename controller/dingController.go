package controller

import (
	"fmt"
	"gin_server_v1/helper"
	"gin_server_v1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendMsg struct {
	Text string `form:"text" json:"text" xml:"text"  binding:"required"`
}

func DingSend(c *gin.Context) {
	var SendMsg SendMsg
	if err := c.ShouldBind(&SendMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "asdasdasd"})
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

func DbText(c *gin.Context) {
	fmt.Println(model.GetById(1))
}
