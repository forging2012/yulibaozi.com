package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

type User struct {
	Id         int64  `json:"id"`
	Portrait   string `json:"portrait"`                     //头像
	Nickname   string `json:"nickname"  xorm:"varchar(11)"` //昵称
	Password   string `json:"password"`                     //密码
	Aword      string `json:"aword"`                        //一句话
	Createtime int64  `json:"createtime" xorm:"createtime"` //注册时间
}

func init() {
	orm.GetEngine().CreateTables(new(User))
}
func (user *User) TableName() string {
	return "user"
}

func (user *User) Inset() (newId int64, err error) {
	engine := orm.GetEngine()

	newId, err = engine.Insert(user)
	return
}

func (user *User) Delete() (delId int64, err error) {
	engine := orm.GetEngine()

	delId, err = engine.Delete(user)
	return
}

func (user *User) Update() (updId int64, err error) {
	engine := orm.GetEngine()

	updId, err = engine.Id(user.Id).Update(user)
	return
}

// ok: false 未找到;true 找到
func (user *User) GetOne(id int64) (ok bool, err error) {
	engine := orm.GetEngine()

	ok, err = engine.Id(id).Get(user)
	return
}

// InsetOrUpdate 插入或者更新，当没有这个数据的时候就插入 如果有就更新
// id作为标记
func (user *User) InsetOrUpdate() (inorUpdId int64, err error) {

	var (
		ok bool
	)
	if user.Id <= 0 {
		//插入
		inorUpdId, err = user.Inset()
		return
	}
	ok, err = new(User).GetOne(user.Id)
	if err == nil && ok {
		//更新
		inorUpdId, err = user.Update()
		return
	}
	// 插入
	inorUpdId, err = user.Inset()
	return
}
