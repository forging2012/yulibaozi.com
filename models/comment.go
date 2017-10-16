package models

import (
	"time"

	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

type Comment struct {
	Id         int64     `json:"id"`
	RowId      string    `json:"rowid"`      //当前行id
	ParentId   string    `json:"parentid"`   //父id
	Aid        int64     `json:"aid"`        //文章id
	ToUserName string    `json:"tousername"` //二级回复时
	NickName   string    `json:"nickname"`
	Email      string    `json:"email"`
	WebSite    string    `json:"website"`
	Content    string    `json:"content"`
	IsView     int8      `json:"isview"` // 0:未审核 1:审核不通过 2:审核通过
	Ip         string    `json:"ip"`     //评论的ip地址
	CreateTime time.Time `json:"createtime"`
}

func init() {
	orm.GetEngine().CreateTables(new(Comment))
}

func (comment *Comment) TableName() string {
	return "comment"
}
