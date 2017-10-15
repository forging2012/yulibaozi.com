package controllers

import (
	"strconv"

	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/service"
)

type ArticleController struct {
	BaseController
}

// GetOne 通过文章id获取文章详情
func (this *ArticleController) GetOne(ctx dotweb.Context) (err error) {
	aid, err := strconv.ParseInt(ctx.QueryString("aid"), 10, 64)
	if aid <= 0 || err != nil {
		this.Respone(ctx, 1, "aid获取出错,请传入数字", nil)
		return
	}
	article, err := new(service.ArticleService).GetOne(aid)
	if err != nil {
		this.Respone(ctx, 1, "获取文章出错", nil)
		return
	}
	err = this.Respone(ctx, 0, "获取文章出错", article)
	return
}

// Page 分页显示文章列表，通过时间倒叙排列
func (this *ArticleController) Page(ctx dotweb.Context) (err error) {
	limit, err := strconv.Atoi(ctx.QueryString("pageSize"))
	if limit <= 0 || err != nil {
		this.Respone(ctx, 1, "未知异常", nil)
		return
	}
	//获取总数
	count, err := new(service.ArticleService).Total()
	if err != nil {
		this.Respone(ctx, 1, "未知异常2", nil)
		return
	}
	pageMtd := this.SetPaginator(limit, count, ctx)
	list, err := new(service.ArticleService).PageArticle(pageMtd.Offset(), limit)
	if err != nil {
		this.Respone(ctx, 1, "未知异常3", nil)
		return

	}
	this.Respone(ctx, 0, "", list)
	return
}

// PageByCid 通过分类id，分页显示文章列表，通过时间倒叙排列
func (this *ArticleController) PageByCid(ctx dotweb.Context) (err error) {
	cid, err := strconv.ParseInt(ctx.QueryString("cid"), 10, 64)
	if err != nil || cid <= 0 {
		this.Respone(ctx, 0, "未知异常1", nil)
		return
	}
	limit, err := strconv.Atoi(ctx.QueryString("pageSize"))
	if limit <= 0 || err != nil {
		this.Respone(ctx, 1, "未知异常", nil)
		return
	}
	//通过分类id获取文章总数
	cat := &service.CategoryService{}
	count, err := cat.GetCount(cid)
	if err != nil {
		this.Respone(ctx, 1, "未知异常2", nil)
		return
	}
	pageMtd := this.SetPaginator(limit, count, ctx)
	list, err := new(service.ArticleService).PageArticleByCid(pageMtd.Offset(), limit, cid)
	if err != nil {
		this.Respone(ctx, 1, "未知异常三", nil)
		return
	}
	this.Respone(ctx, 0, "", list)
	return

}

// Like 猜我喜欢，通过当前文章id，查询到分类id，然后通过分类id，查询三篇随机文章
func (this *ArticleController) Like(ctx dotweb.Context) (err error) {
	cid, err := strconv.ParseInt(ctx.QueryString("cid"), 10, 64)
	if err != nil || cid <= 0 {
		this.Respone(ctx, 1, "未知异常1", nil)
		return
	}
	list, err := new(service.ArticleService).ArticleLike(cid)
	if err != nil {
		this.Respone(ctx, 1, "未知异常1", nil)
		return
	}
	this.Respone(ctx, 0, "", list)
	return

}

//Hot 热门文章，通过浏览数倒叙排序 列表，并限制一定的大小
func (this *ArticleController) Hot(ctx dotweb.Context) (err error) {
	list, err := new(service.ArticleService).Hot()
	if err != nil {
		this.Respone(ctx, 1, "未知异常1", nil)
		return
	}
	this.Respone(ctx, 0, "", list)
	return
}
