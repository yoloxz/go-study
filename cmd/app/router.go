package app

import (
	"go-study/cmd/api"
	"go-study/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title IvanApi Swagger 标题
// @version 1.0 版本
// @description IvanApi Service 描述
// @BasePath /api/v1  基础路径
// @query.collection.format multi
func RouterInit() {

	// 创建一个 gin 实例
	router := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	url := ginSwagger.URL("http://localhost:3333/swagger/doc.json")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userRouter := router.Group("user")

	userRouter.GET("/userInfo", api.GetUserInfo)

	userRouter.POST("/register", api.Register)

	fileRouter := router.Group("file")

	fileRouter.POST("/uploadFile", api.UploadFile)
	fileRouter.GET("/downloadFile/:filename", api.DownloadFile)
	fileRouter.GET("/getMd5", api.GetMd5)

	// 解决方案路由组
	solutionRouter := router.Group("solution")
	solutionRouter.GET("/list", api.GetSolutionList)
	solutionRouter.POST("/add", api.AddSolution)
	solutionRouter.POST("/update", api.UpdateSolution)
	solutionRouter.POST("/delete", api.DeleteSolution)

	router.Run(":3333")

}
