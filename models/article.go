package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/initialization"
)

// Article 文章模型
type Article struct {
	Id           int64  `json:"id"`
	Userid       int64  `json:"userid"`  //作者id
	Author       string `json:"author"`  //作者名字
	CId          int64  `json:"cid"`     //标签id
	Seotitle     string `json"seotitle"` //seo的标题
	Picture      string `json:"picture"`
	Keywords     string `json:"keywords"`                //seo的关键字 建议个数在5-10之间，用英文逗号隔开
	Description  string `json:"description"`             //seo描述  建议字数在30-70之间
	Title        string `json:"title"`                   //标题
	Content      string `json:"content" `                //内容
	Thumbscount  int    `json:"thumbscount"`             //点赞数
	Viewcount    int    `json:"viewcount"`               //阅读次数
	Commentcount int    `json:"commentcount"`            //评论次数
	Updatedat    int64  `json:"updateat" xorm:"updated"` //更新时间
	Year         int    `json:"year"`                    //发布的年
	Month        int    `json:"month"`                   //发布的日期
	Day          int    `json:"day"`                     //发布的天
	ReleaseTime  int64  `json:"releasetime"`             //发布时间
	Copyright    string `json:"Copyright"`               //文章底部版权
}
func init() {
	orm.GetEngine().Table(new(Article))
}

func (article *Article) TableName() string {
	return "article"
}

func (article *Article) Inset() (newId int64, err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	newId, err = engine.Insert(article)
	return
}

func (article *Article) Delete() (delId int64, err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	delId, err = engine.Delete(article)
	return
}

func (article *Article) Update() (updId int64, err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	updId, err = engine.Id(article.Id).Update(article)
	return
}

// UpdateViewCount 专门更新viewCount
func (article *Article) UpdateViewCount() (err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	_, err = engine.Id(article.Id).Update(article)
	return
}

func (article *Article) GetOne(id int64) (ok bool, err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	ok, err = engine.Id(id).Get(article)
	return
}

// TopN 获取文章列表
func (article *Article) TopN(n int) (articles []*Article, err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	err = engine.Desc("id").Limit(n).Find(&articles)
	return
}

// PageUser 分页的文章数
func (article *Article) PageArticle(offset, limit int) (articles []*Article, err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	err = engine.Limit(int(limit), int(offset)).Find(&articles)
	return
}

// GetArticleByTag 通过标签列表
// func (article *Article) GetArticleByTag(ctid int64) (articles []*Article, err error) {
// 	// err=engine.Join("INNER","tag_article_rel","tag_article_rel").Find(&articles)
// 	err = engine.Where("Ctid =?", ctid).Join("INNER", "tag_article_rel", "tag_article_rel.cid=article.id").Desc("id").Find(&articles)

// 	return
// }

// GetArticleByTag 通过标签列表
// func (article *Article) PageArticleByTag(ctid, offset, limit int64) (articles []*Article, err error) {
// 	// err=engine.Join("INNER","tag_article_rel","tag_article_rel").Find(&articles)
// 	err = engine.Where("Ctid =?", ctid).Join("INNER", "tag_article_rel", "tag_article_rel.cid=article.id").Desc("id").Limit(int(limit), int(offset)).Find(&articles)
// 	return
// }

// // Count 统计所有文章的数量
func (article *Article) Total() (count int64, err error) {
	engine := orm.GetEngine()
	defer engine.Close()
	count, err = engine.Count(article)
	return
}

//通过
