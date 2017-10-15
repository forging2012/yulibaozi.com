package models
import (
	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

// EmailInfo 发送邮件的信息
type EmailInfo struct {
	Id           int64  `json:"id"`
	Aid          int64  `json:"aid"`          //文章id
	Sendname     string `json:"sendname"`     //发送人昵称
	Sendemail    string `json:"sendemial"`    //发送人的邮箱
	Website      string `json:"website"`      //发送人的站点
	Permissions  int    `json:"permissions"`  //展示权限
	Emailcontent string `json:"emailcontent"` //发送内容
	Clientip     string `json:"clientip"`     //客户端ip

}

func (mail *EmailInfo) Inset() (newId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	newId, err = engine.Insert(mail)
	return
}

func (mail *EmailInfo) Delete() (delId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	delId, err = engine.Delete(mail)
	return
}

func (mail *EmailInfo) Update() (updId int64, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	updId, err = engine.Id(mail.Id).Update(mail)
	return
}

func (mail *EmailInfo) GetOne(id int64) (ok bool, err error) {
	engine:=orm.GetEngine()
	defer engine.Close()
	ok, err = engine.Id(id).Get(mail)
	return
}
