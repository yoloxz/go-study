package model

import "time"

type StoreSolution struct {
	Id           int       `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string    `gorm:"type:varchar(100);not null"`
	Version      int       `gorm:"type:int;not null"`
	Type         int       `gorm:"type:int;not null"`
	Introduction string    `json:"introduction" form:"introduction"`
	PicUrl       string    `json:"picUrl" form:"picUrl"`
	Status       int       `json:"status" form:"status"`
	CreateTime   time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	Updator      string    `gorm:"type:varchar(100)"`
	Extend1      string    `gorm:"type:varchar(100)"`
	Extend4      string    `gorm:"type:varchar(100)"`
	Extend2      string    `gorm:"type:varchar(100)"`
	Extend3      string    `gorm:"type:varchar(100)"`
	Extend5      string    `gorm:"type:varchar(100)"`
	Creator      string    `gorm:"type:varchar(100)"`
}

type AddSolutionReq struct {
	Name         string `json:"name" form:"name"`
	Version      int    `json:"version" form:"version"`
	Type         int    `json:"type" form:"type"`
	Introduction string `json:"introduction" form:"introduction"`
	PicUrl       string `json:"picUrl" form:"picUrl"`
	Status       int    `json:"status" form:"status"`
}
