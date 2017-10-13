package controllers

import "github.com/devfeel/dotweb"
import "fmt"

type FileController struct {
	BaseController
}

func (filec *FileController) Insert(ctx dotweb.Context) (err error) {
	file, err := ctx.Request().FormFile("file")
	if err != nil {
		ctx.WriteString("上传文件出错")
		filec.Respone(ctx, -1, "上传文件错误", nil)
		return
	}
	var size int64
	fmt.Println("file.FileName():",file.FileName())
	size, err = file.SaveFile("/Users/yulibaozi/GoWorkSpace/src/yulibaozi/file/" + fmt.Sprintf("%s:%s","hello",file.FileName()))
	if err != nil {
		ctx.WriteString("保存文件出错")
		return
	}
	ctx.WriteString(fmt.Sprintf("保存文件成功:%d", size))
	return

}

func (filec *FileController) Show(ctx dotweb.Context) (err error) {
	err = ctx.View("/Users/yulibaozi/GoWorkSpace/src/yulibaozi/views/index.html")
	return
}
