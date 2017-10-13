package json

import (
	"encoding/json"
)

// JsonToString 把json对象转换为字符串
// obj json对象
func JsonToString(obj interface{}) (jsonStr string,err error)  {
	var (
		bytes []byte
	)
	bytes,err= json.Marshal(obj)
	if err != nil {
		return
	}
	jsonStr = string(bytes)
	return
}
// StringToJson 把字符串转换为json对象
// jsonStr json字符串
// obj 接收json对象的容器
func StringToJson(jsonStr string,obj interface{}) (err error)  {
	err= json.Unmarshal([]byte(jsonStr),obj)
	return
	
}


