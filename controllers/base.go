package controllers

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/devfeel/dotweb"
)

type BaseController struct {
}

// DecodeJSONReq 从请求体中解析json对象
// obj 解析到的对象
func (this *BaseController) DecodeJSONReq(ctx dotweb.Context, obj interface{}) {

	err := ctx.Bind(obj)
	if err != nil {
		this.Respone(ctx, 1, err.Error(), nil)
		return
	}
}

// CreateToken 生成token,作用重复提交表单
func (this *BaseController) CreateToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

// IsSubmitAgain 是否是重复提交
// token 传入的token
// ok:false 不是重复提交;true 是重复提交
func (this *BaseController) IsSubmitAgain(ctx dotweb.Context, token string) bool {
	cotoken, err := ctx.ReadCookieValue("token")
	if err != nil {
		return true
	}
	if token == "" || len(token) == 0 || token != cotoken || strings.Compare(cotoken, token) != 0 {
		return true
	}
	return false
}

// DataResponse 响应结构体
type DataResponse struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	ServerTime int64       `json:"stime"`
	Body       interface{} `json:"body"`
}

//Respone json输出
// errCode 自定义业务返回码
// msg 返回结果
// body 返回数据
func (this *BaseController) Respone(ctx dotweb.Context, errCode int, msg string, body interface{}) (err error) {
	resp := &DataResponse{
		Code:       errCode,
		Msg:        msg,
		ServerTime: time.Now().Unix(),
		Body:       body,
	}
	_, err = ctx.WriteJson(resp)
	return
}
