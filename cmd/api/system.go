package api

import (
	"github.com/gin-gonic/gin"
)

// @Summary 获取系统用户列表
// @Description 获取系统用户列表
// @Tags 系统
// @Accept json
// @Produce json
// @Success 200 {object} string "ok"
// @Router /system/list [get]
func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user list",
	})
}
