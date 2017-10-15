package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

type SmtpConfig struct {
	Id            int64  `json:"id"`
	Smtp          string `json:"smtp" xorm:"varchar(11)"`         //smtp服务器  例如，smtp.163.com（163邮箱）， smtp.qq.com（QQ邮箱）
	Smtpsecurity  string `json:"smtpsecurity" xorm:"varchar(11)"` //smtp的加密方式
	Smtpport      int    `json:"smtpport"`                        //smtp的端口号
	Emailaddress   string `json:"emailaddress" xorm:"varchar(11)"`  //邮箱地址
	Smtppwd       string `json:"smtppwd"`                         //smtp密码
	Emailtemplate string `json:"emailtemp"`                       //邮件模板
}

func (smtpConf *SmtpConfig) Inset() (newId int64, err error) {
	engine:=orm.GetEngine()
	
	newId, err = engine.Insert(smtpConf)
	return
}

func (smtpConf *SmtpConfig) Delete() (delId int64, err error) {
	engine:=orm.GetEngine()
	
	delId, err = engine.Delete(smtpConf)
	return
}

func (smtpConf *SmtpConfig) Update() (updId int64, err error) {
	engine:=orm.GetEngine()
	
	updId, err = engine.Id(smtpConf.Id).Update(smtpConf)
	return
}

func (smtpConf *SmtpConfig) GetOne(id int64) (ok bool, err error) {
	engine:=orm.GetEngine()
	
	ok, err = engine.Id(id).Get(smtpConf)
	return
}
