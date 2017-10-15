package controllers

import (
	"strconv"

	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/service"
)

type CatController struct {
	BaseController
}

// GetList 获取分页列表，通过列表数排列
func (this *CatController) GetList(ctx dotweb.Context) (err error) {
	limit, err := strconv.Atoi(ctx.QueryString("pageSize"))
	if limit <= 0 || err != nil {
		this.Respone(ctx, 1, "未知异常", nil)
		return
	}
	count, err := new(service.CategoryService).GetCount()
	if err != nil {
		this.Respone(ctx, 1, "未知异常1", nil)
		return
	}
	pageMtd := this.SetPaginator(limit, count, ctx)
	list, err := new(service.CategoryService).PageCats(pageMtd.Offset(), limit)
	if err != nil {
		this.Respone(ctx, 1, "未知异常2", nil)
		return
	}
	this.Respone(ctx, 0, "", list)
	return
}
