package api

import (
	"go-study/internal/app"
	"go-study/internal/model"
	"go-study/internal/model/request"

	"go-study/internal/service"

	"go-study/global"

	"github.com/gin-gonic/gin"
)

// var logger = global.LOGGER
// var reqLogger = app.NewReqLogger

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param name query string true "用户名"
// @Success 200 {object} string "ok"
// @Router /user/userInfo [get]
func GetUserInfo(c *gin.Context) {

	repUser, err := service.GetUserById(1)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"userInfo": repUser,
	})
}

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param name body string true "用户名"
// @Param age body int true "年龄"
// @Param email body string true "邮箱"
// @Param password body string true "密码"
// @Success 200 {object} string "ok"
// @Router /user/register [post]
func Register(c *gin.Context) {

	logger := global.LOGGER

	app.NewReqLogger(c)

	var u request.Register

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		logger.Error().Msgf("参数错误: %v", err)
		return
	}

	user := &model.User{Name: u.Name, Age: u.Age, Email: u.Email, PassWord: u.PassWord}

	resUser, err := service.Register(*user)

	if err != nil {
		logger.Error().Msgf("注册失败: %v", err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"userInfo": resUser,
	})
}
