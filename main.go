package main

import (
	"go-study/internal/initialize"
	"os"
)

func main() {

	// 初始化cobra
	rootCmd := initialize.NewRootCommnad()
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}

	//sqlite创建test表
	// initialize.GromSqliteInit("./test.db").AutoMigrate(&model.StoreSolution{})
}
