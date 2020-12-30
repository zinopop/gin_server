package v6

import (
	"gin_server/model/andoncloud"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DingController struct {
}

func (DingController) Test(c *gin.Context) {
	//BPId,err := strconv.ParseInt(c.Request.FormValue("BPId"), 10, 64)
	params := struct {
		BPId int64 `form:"BPId" json:"BPId" xml:"BPId"  binding:"required"`
	}{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := andoncloud.Bloodpressureinfo{}.FindOne(params.BPId)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"data":    "",
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    data,
		"message": "success",
	})
}
