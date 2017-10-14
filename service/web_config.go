package service

import (
	"errors"
	"github.com/yulibaozi/yulibaozi.com/models"
)

type WebConfService struct{}

func (webconfService *WebConfService) Insert(webConf *models.WebConfig) (newId int64, err error) {
	newId, err = webConf.Inset()
	return
}

func (webconfService *WebConfService) Update(webConf *models.WebConfig) (updId int64, err error) {
	updId, err = webConf.Update()
	return
}

func (webconfService *WebConfService) Delete(webConf *models.WebConfig) (delId int64, err error) {
	delId, err = webConf.Delete()
	return
}

func (webconfService *WebConfService) GetUser(uid int64) (webConf *models.WebConfig, err error) {
	var (
		ok bool
	)
	ok, err = webConf.GetOne(uid)
	if !ok {
		errors.New("没有找到数据库")
		return
	}
	return
}