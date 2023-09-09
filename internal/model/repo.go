package model

import (
	"sync"

	"gorm.io/gorm"
)

// 全局数据库连接
type InstallRepo struct {
	*sync.RWMutex
	*gorm.DB
}

var RepoDb *InstallRepo
