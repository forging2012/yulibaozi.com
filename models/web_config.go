package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

type WebConfig struct {
	Id              int64  `json:"id"`
	Homekeywords    string `json:"homekeywords"`    //关键字有利于SEO优化，建议个数在5-10之间，用英文逗号隔开
	Homedescription string `json:"homedescription"` //首页描述 ->描述有利于SEO优化，建议字数在30-70之间
	Webicon         string `json:"webicon"`         //网站icon地址
	Themecolor      string `json:"themecolor"`      //主题颜色  颜色选项，和哀悼
	Footinfo        string `json:"footinfo"`        //底部信息
}

func (webConf *WebConfig) Inset() (newId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	newId, err = engine.Insert(webConf)
	return
}

func (webConf *WebConfig) Delete() (delId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	delId, err = engine.Delete(webConf)
	return
}

func (webConf *WebConfig) Update() (updId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	updId, err = engine.Id(webConf.Id).Update(webConf)
	return
}

func (webConf *WebConfig) GetOne(id int64) (ok bool, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	ok, err = engine.Id(id).Get(webConf)
	return
}
