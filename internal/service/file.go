package service

import (
	"crypto/md5"
	"fmt"
	"go-study/global"
	"go-study/internal/app"
	"go-study/internal/model"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	var reqLogger = app.NewReqLogger
	var logger = global.LOGGER

	reqLogger(c)
	// 解析表单字段和文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(400, "Bad Request")
		reqLogger(c)
		return
	}

	// 创建目标文件
	dstPath := "/home/sike/uploads/" + file.Filename

	//判断数据库filename是否重复
	var fileDB model.File
	global.DB.Where("filename = ?", file.Filename).First(&fileDB)
	if fileDB.ID != 0 {
		c.String(400, "文件已存在")
		logger.Err(err).Msg("文件已存在")
		return
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		log.Println(err)
		c.String(500, "Internal Server Error")
		logger.Err(err).Msg("文件上传失败")
		return
	}
	defer dstFile.Close()

	// 打开上传的文件
	srcFile, err := file.Open()
	if err != nil {
		log.Println(err)
		c.String(500, "Internal Server Error")
		return
	}
	defer srcFile.Close()

	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		logger.Err(err).Msg("文件上传失败")
		c.String(500, "Internal Server Error")
		return
	}

	// 文件路径保存到数据库
	saveFilePathToDB(file.Filename, dstPath)

	c.String(200, "File uploaded successfully!")
}

// 获取文件列表
func GetFileList() (fileList []model.File, err error) {
	db := global.DB
	err = db.Find(&fileList).Error
	return fileList, err
}

// 下载文件
func DownloadFile(c *gin.Context) {
	var reqLogger = app.NewReqLogger
	var logger = global.LOGGER

	reqLogger(c)
	// 获取文件名
	filename := c.Param("filename")

	// 获取文件路径
	var file model.File
	global.DB.Where("filename = ?", filename).First(&file)
	filepath := file.Filepath

	// 打开文件
	f, err := os.Open(filepath)
	if err != nil {
		logger.Err(err).Msg("文件打开失败")
		c.String(500, "Internal Server Error")
		return
	}
	defer f.Close()

	// 设置响应头
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filename)

	// 读取文件内容并写入响应
	_, err = io.Copy(c.Writer, f)
	if err != nil {
		logger.Err(err).Msg("文件读取失败")
		c.String(500, "Internal Server Error")
		return
	}
}

func GetMd5(c *gin.Context) (string, error){
	// 利用/proc/self/exe获取可执行文件

	file, err := os.Open("/proc/self/exe")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return "", err
	}
	defer file.Close()

	exePath, err := os.Readlink("/proc/self/exe")
	fmt.Println(exePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 创建一个 MD5 hasher
	hasher := md5.New()

	// 将文件内容复制到 hasher 中
	if _, err := io.Copy(hasher, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return "", err
	}

	// 计算 MD5 值
	sum := hasher.Sum(nil)

	return fmt.Sprintf("%x", sum), nil
}

func saveFilePathToDB(filename string, path string) {
	db := global.DB
	var file model.File
	// 获取文件名
	file.Filename = filename
	file.Filepath = path

	db.Create(&file)
}

// 检查可执行文件的MD5
