package global

import (
	"time"
)

type MODEL struct {
	ID         uint      `gorm:"primarykey"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
