package controllers

import (
	"fmt"
	"strconv"
	"github.com/yulibaozi/yulibaozi.com/models"

	"github.com/devfeel/dotweb"
)

type UserController struct {
	BaseController
}

func (this *UserController) InsetOrUpdate(ctx dotweb.Context) (err error) {
	var (
		newId int64
	)
	user := &models.User{}
	this.DecodeJSONReq(ctx, user)
	fmt.Println("USER:", user)
	newId, err = user.InsetOrUpdate()
	if err != nil {
		this.Respone(ctx, 1, err.Error(), nil)
		return
	}
	this.Respone(ctx, 0, "请求成功", newId)
	return
}

func (this *UserController) Delete(ctx dotweb.Context) (err error) {
	var (
		id int64
	)
	
	uid, err := strconv.ParseInt(ctx.QueryString("uid"), 10, 64)
	if uid <= 0 || err != nil {
		this.Respone(ctx, 1, "uid获取出错,请传入数字", nil)
		return
	}
	user := &models.User{}
	ok, _ := user.GetOne(uid)
	if !ok {
		this.Respone(ctx, 1, "数据不存在", nil)
		return
	}
	id, err = user.Delete()
	if err != nil {
		this.Respone(ctx, -1, "删除失败", nil)
		return
	}
	this.Respone(ctx, 0, "删除成功", id)
	return
}

func (this *UserController) GetOne(ctx dotweb.Context) (err error) {
	var (
		ok bool
	)
	uid, err := strconv.ParseInt(ctx.QueryString("uid"), 10, 64)
	if uid <= 0 || err != nil {
		this.Respone(ctx, 1, "uid获取出错,请传入数字", nil)
		return
	}
	user := &models.User{}
	ok, err = user.GetOne(uid)
	if err != nil {
		this.Respone(ctx, -1, "获取对象失败:"+err.Error(), nil)
		return
	}
	if !ok {
		this.Respone(ctx, -1, "不存在该数据", nil)
		return
	}
	this.Respone(ctx, 0, "获取成功", user)
	return
}
