package api

import (
	"go-study/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary 上传文件
// @Description 上传文件
// @Tags 文件
// @Accept json
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} string "ok"
// @Router /file/upload [post]
func UploadFile(c *gin.Context) {
	service.UploadFile(c)
}

// @Summary 下载文件
// @Description 下载文件
// @Tags 文件
// @Accept form-data
// @Produce json
// @Param filename query string true "文件名"
// @Success 200 {object} string "ok"
// @Router /file/download [get]
func DownloadFile(c *gin.Context) {
	service.DownloadFile(c)
}

func GetMd5(c *gin.Context) {
	md5str, err := service.GetMd5(c)
	if err != nil {
		c.String(500, err.Error())
		log.Fatal(err)
	}
	c.String(200, md5str)
}
