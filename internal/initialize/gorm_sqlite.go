package initialize

import (
	"go-study/global"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GromSqliteInit(dbFile string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(global.CONFIG.SQLITE.MAX_IDLE_CONNS)
	sqlDB.SetMaxOpenConns(100)

	return db
}
