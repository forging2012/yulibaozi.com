package models


import (
	"time"
	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

type Picture struct {
	Id       int64     `json:"id"`       //id
	Adress   string    `json:"adress"`   //图片地址
	CreateAt time.Time `json:"createat"` //创建时间
}


func (picture *Picture) Inset() (newId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	newId, err = engine.Insert(picture)
	return
}

func (picture *Picture) Delete() (delId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	delId, err = engine.Delete(picture)
	return
}

func (picture *Picture) Update() (updId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	updId, err = engine.Id(picture.Id).Update(picture)
	return
}

func (picture *Picture) GetOne(id int64) (ok bool, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	ok, err = engine.Id(id).Get(picture)
	return
}
