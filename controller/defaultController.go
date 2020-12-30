package controller

import (
	"gin_server/model"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (DefaultController) Index(c *gin.Context) {
	//id,err := model.UserTest{}.NewUser("asdasd")

	userInfo, err := model.UserTest{}.GetUserById(1)

	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"data":    "",
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    userInfo,
		"message": "success",
	})
}
