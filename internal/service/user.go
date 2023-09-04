package service

import (
	"errors"
	"go-study/global"
	"go-study/internal/model"
	"go-study/utils"

	"gorm.io/gorm"
)

func GetUserById(id uint) (user model.User, err error) {

	var reqUser model.User

	err = global.DB.First(&reqUser, "id = ?", id).Error
	if err != nil {
		return user, err
	}
	return reqUser, err
}

func Register(u model.User) (resUser model.User, err error) {
	var user model.User
	if !errors.Is(global.DB.Where("name = ?", u.Name).First(&user).Error, gorm.ErrRecordNotFound) {
		return resUser, errors.New("用户名已存在")
	}
	u.PassWord = utils.BcryptHash(u.PassWord)
	err = global.DB.Create(&u).Error
	return u, err
}
