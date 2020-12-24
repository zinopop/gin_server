package controller

import (
	"github.com/gin-gonic/gin"
	"os/exec"
)

func Rebuild(c *gin.Context) {
	command := "/home/admin/gin_server_v1/build.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		c.JSON(500, gin.H{
			"data":    "",
			"message": "命令行错误",
		})
		return
	}
	resp := string(bytes)
	c.JSON(200, gin.H{
		"data":    "",
		"message": resp,
	})
}