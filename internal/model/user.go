package model

import (
	"go-study/global"
)

type User struct {
	global.MODEL
	Name     string `json:"name" gorm:"column:name"`
	Age      int    `json:"age" gorm:"column:age"`
	Email    string `json:"eamil" gorm:"column:email"`
	PassWord string `json:"-" gorm:"column:password"`
}

func (u *User) TableName() string {
	return "users"
}
