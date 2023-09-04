package model

import (
	"go-study/global"
)

type File struct {
	global.MODEL
	Filename string `json:"filename" gorm:"column:filename"`
	Filepath string `json:"filepath" gorm:"column:filepath"`
}
