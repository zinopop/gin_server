package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (*BaseController) Test() string {
	return "aaaaa"
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (*BaseController) SuccessJsonResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func (*BaseController) ErrorJsonResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, response{
		Code:    500,
		Message: message,
		Data:    data,
	})
}
