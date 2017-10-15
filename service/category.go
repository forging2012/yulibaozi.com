package service

import (
	"fmt"

	"github.com/yulibaozi/yulibaozi.com/models"
)

type CategoryService struct{}

// Get 通过分类id获取文章总数
func (categoryService *CategoryService) GetCount(cid int64) (count int64, err error) {
	var ok bool
	cat := &models.Category{}
	ok, err = cat.GetOne(cid)
	if err != nil {
		return
	}
	if !ok {
		err = fmt.Errorf("获取文章总数失败")
		return
	}
	count = cat.Count
	return
}
