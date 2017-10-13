package service

import (
	"errors"
	"yulibaozi/models"
)

type UserService struct{}

func (userService *UserService) Insert(user *models.User) (newId int64, err error) {
	newId, err = user.Inset()
	return
}

func (userService *UserService) Update(user *models.User) (updId int64, err error) {
	updId, err = user.Update()
	return
}

func (userService *UserService) Delete(user *models.User) (delId int64, err error) {
	delId, err = user.Delete()
	return
}

func (userService *UserService) GetUser(uid int64) (user *models.User, err error) {
	var (
		ok bool
	)
	ok, err = user.GetOne(uid)
	if !ok {
		errors.New("没有找到数据库")
		return
	}
	return
}
