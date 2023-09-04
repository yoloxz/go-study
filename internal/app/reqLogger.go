package app

import (
	"fmt"
	"go-study/global"

	"github.com/gin-gonic/gin"
)

func NewReqLogger(c *gin.Context) {
	logger := global.LOGGER
	fmt.Println(c)
	// 请求消息
	logger.Info().
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("ip", c.ClientIP()).
		Msg("请求消息")
}
