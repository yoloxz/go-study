package global

import (
	"go-study/config"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var (
	// 全局配置
	DB     *gorm.DB
	CONFIG config.CONFIG
	LOGGER zerolog.Logger
)
