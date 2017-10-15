package models

// 文章分类和标签
type Category struct {
	Id    int64 `json:"id"`
	Kind  int8  `json:"kind"`  //0：标签 1: 分类
	Count int64 `json:"count"` //文章总数
}

func (cat *Category) Inset() (newId int64, err error) {
	newId, err = engine.Insert(cat)
	return
}

func (cat *Category) Delete() (delId int64, err error) {
	delId, err = engine.Delete(cat)
	return
}

func (cat *Category) Update() (updId int64, err error) {
	updId, err = engine.Id(cat.Id).Update(cat)
	return
}

func (cat *Category) GetOne(id int64) (ok bool, err error) {
	ok, err = engine.Id(id).Get(cat)
	return
}

func (cat *Category) Total() (count int64, err error) {
	count, err = engine.Count(cat)
	return
}

func (cat *Category) PageCats(offset, limit int) (cats []*Category, err error) {
	err = engine.Limit(limit, offset).Find(&cats)
	return
}
