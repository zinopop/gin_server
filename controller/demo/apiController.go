package demo

import (
	"gin_server/controller"
	"gin_server/model"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	controller.BaseController
}

func (controller ApiController) GetTest(c *gin.Context) {
	userInfo, err := model.UserTest{}.GetUserById(1)
	if err != nil {
		controller.ErrorJsonResponse(c, err.Error(), "")
		return
	}
	controller.SuccessJsonResponse(c, "获取成功", userInfo)
}
