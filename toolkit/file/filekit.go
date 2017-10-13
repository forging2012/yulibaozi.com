package file

import (
	"os"
	"strings"
)

//GetFileExt 获取文件的扩展名
// fileName:文件名
// ext 扩展名
func GetFileExt(fileName string) (ext string) {
	if len(fileName) == 0 || fileName == "" {
		return
	}
	index := strings.LastIndex(fileName, ".")
	if index < 0 {
		return
	}
	ext = string(fileName[index:])
	return
}

//Exist 判断文件名是否存在
// fileName:文件名

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
