package v6

import (
	"gin_server/controller"
	"gin_server/model/andoncloud"
	"github.com/gin-gonic/gin"
)

type DingController struct {
	controller.BaseController
}

func (controller DingController) Test(c *gin.Context) {
	//BPId,err := strconv.ParseInt(c.Request.FormValue("BPId"), 10, 64)
	//参数绑定和入参
	params := struct {
		BPId int64 `form:"BPId" json:"BPId" xml:"BPId"  binding:"required"`
	}{}
	//入参校验
	if err := c.ShouldBind(&params); err != nil {
		controller.ErrorJsonResponse(c, err.Error(), "")
		return
	}
	//调用模型层获取持久层数据
	data, err := andoncloud.Bloodpressureinfo{}.FindOne(params.BPId)
	//处理错误
	if err != nil {
		controller.ErrorJsonResponse(c, err.Error(), "")
		return
	}
	//相应参数
	controller.SuccessJsonResponse(c, "获取成功", data)
}
