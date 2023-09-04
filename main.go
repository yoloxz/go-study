package main

import (
	"go-study/cmd/app"
	"go-study/config"
	"go-study/global"
	"go-study/internal/initialize"
)

func main() {
	// 加载配置
	global.CONFIG = *config.NewConfig("./config.yaml")

	// 日志初始化
	global.LOGGER = initialize.LoggerInit()

	// sqlite初始化
	global.DB = initialize.GromSqliteInit("./test.db")

	// 启动路由
	app.RouterInit()
}
