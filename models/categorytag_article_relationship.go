package models

import (
	"fmt"
)

// TagArticleRel 文章和标签表
type TagArticleRel struct {
	Id   int64 `json:"id"`
	Aid  int64 `json:"aid"`  //文章id
	Ctid int64 `json:"ctid"` //标签或者分类id
}

func (tagArticleRel *TagArticleRel) Inset() (newId int64, err error) {
	newId, err = engine.Insert(tagArticleRel)
	return
}

func (tagArticleRel *TagArticleRel) Delete() (delId int64, err error) {
	delId, err = engine.Delete(tagArticleRel)
	return
}

func (tagArticleRel *TagArticleRel) Update() (updId int64, err error) {
	updId, err = engine.Id(tagArticleRel.Id).Update(tagArticleRel)
	return
}

func (tagArticleRel *TagArticleRel) GetOne(id int64) (ok bool, err error) {
	ok, err = engine.Id(id).Get(tagArticleRel)
	return
}

func (tagArticleRel *TagArticleRel) CountByCtid(ctid int64) (count int64,err error) {
	tagArticleRel.Ctid=ctid
	count,err= engine.Count(tagArticleRel)
	return
}

// GetArticleByTag2 这种方式对排序不是很方便
func (tagArticleRel *TagArticleRel) GetArticleByTag2(ctid int64)(articles []*Article, err error)  {
	rels:=make([]*TagArticleRel, 0)
	article:=&Article{}
	err=engine.Where("Ctid =?",ctid).Find(&rels)
	if err != nil {
		return
	}
	if len(rels)<=0 {
		err=fmt.Errorf("未查询到相关数据")
		return
	}
	for _,v := range rels {
		_,err= engine.Id(v.Ctid).Get(article)
		if err != nil {
			return
		}else{
			articles=append(articles,article)
			return
		}
	}
	return
}
