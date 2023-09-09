package initialize

import (
	"go-study/cmd/app"
	"go-study/config"
	"go-study/global"

	"github.com/spf13/cobra"
)

var (
	cfgFile = "./config.yaml"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "go-study version",
		Long:  `All software has versions. This is go-study's`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("go-study v1.0.0")
		},
	}
)

func NewRootCommnad() *cobra.Command {
	rootCommand := &cobra.Command{
		Use: "go-study",
		Long: `useage: go-study [options]
		-h, --help		显示帮助
		-v, --version	显示版本号
		config [filePath]		配置文件路径`,
		Run: func(cmd *cobra.Command, args []string) {
			// 加载配置
			global.CONFIG = *config.NewConfig(cfgFile)

			// 日志初始化
			global.LOGGER = LoggerInit()

			// sqlite初始化
			global.DB = GromSqliteInit(global.CONFIG.SQLITE.DB_FILE_PATH)

			// 启动路由
			app.RouterInit()
		},
	}

	rootCommand.PersistentFlags().StringVar(&cfgFile, "config", "./config.yaml", "config file (default is ./config.yaml)")
	rootCommand.AddCommand(versionCmd)

	return rootCommand
}
