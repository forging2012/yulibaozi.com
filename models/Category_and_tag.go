package models


// 文章分类和标签
type CategoryAndTag struct{
	Id int64  `json:"id"`
	Kind int8 `json:"kind"`  //0：标签 1: 分类
	Count int64 `json:"count"` //文章总数
}

func (cat *CategoryAndTag) Inset() (newId int64, err error) {
	newId, err = engine.Insert(cat)
	return
}

func (cat *CategoryAndTag) Delete() (delId int64, err error) {
	delId, err = engine.Delete(cat)
	return
}

func (cat *CategoryAndTag) Update() (updId int64, err error) {
	updId, err = engine.Id(cat.Id).Update(cat)
	return
}

func (cat *CategoryAndTag) GetOne(id int64) (ok bool, err error) {
	ok, err = engine.Id(id).Get(cat)
	return
}
func (cat *CategoryAndTag) GetCount(id int64) (count int64,err error) {
	err=engine.Id(id).Find(cat)
	if err != nil {
		return
	}
	count=cat.Count
	return
}