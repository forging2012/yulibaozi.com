package service

import (
	"errors"
	"yulibaozi/models"
)

type PictureService struct{}

func (pictureService *PictureService) Update(picture *models.Picture) (updId int64, err error) {
	updId, err = picture.Update()
	return
}

func (pictureService *PictureService) Inset(picture *models.Picture) (newId int64, err error) {
	newId, err = picture.Inset()
	return
}

func (pictureService *PictureService) Delete(picture *models.Picture) (delId int64, err error) {
	delId, err = picture.Delete()
	return
}

func (pictureService *PictureService) GetPicture(uid int64) (picture *models.Picture, err error) {
	var (
		ok bool
	)
	ok, err = picture.GetOne(uid)
	if !ok {
		errors.New("没有找到数据库")
		return
	}
	return
}
