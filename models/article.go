package models

import "time"

// Article 文章模型
type Article struct {
	Id           int64     `json:"id"`
	Userid       int64     `json:"userid"`                  //发送人
	Seotitle     string    `json"seotitle"`                 //seo的标题
	Keywords     string    `json:"keywords"`                //seo的关键字 建议个数在5-10之间，用英文逗号隔开
	Description  string    `json:"description"`             //seo描述  建议字数在30-70之间
	Title        string    `json:"title"`                   //标题
	Content      string    `json:"content" `                //内容
	Viewcount    int       `json:"viewcount"`               //阅读次数
	Commentcount int       `json:"commentcount"`            //评论次数
	Updatedat    time.Time `json:"updateat" xorm:"updated"` //更新时间
	Copyright    string    `json:"Copyright"`               //文章底部版权
}

func (article *Article) Inset() (newId int64, err error) {
	engine:=GetEngine()
	defer engine.Close()
	newId, err = engine.Insert(article)
	return
}

func (article *Article) Delete() (delId int64, err error) {
	delId, err = engine.Delete(article)
	return
}

func (article *Article) Update() (updId int64, err error) {
	updId, err = engine.Id(article.Id).Update(article)
	return
}

func (article *Article) GetOne(id int64) (ok bool, err error) {
	ok, err = engine.Id(id).Get(article)
	return
}

// TopN 获取文章列表
func (article *Article) TopN(n int) (articles []*Article, err error) {
	err = engine.Desc("id").Limit(n).Find(&articles)
	return
}

// PageUser 分页的文章数
func (article *Article) PageArticle(offset, limit int64) (articles []*Article, err error) {
	err = engine.Limit(int(limit), int(offset)).Find(&articles)
	return
}

// GetArticleByTag 通过标签列表
func (article *Article) GetArticleByTag(ctid int64) (articles []*Article, err error) {
	// err=engine.Join("INNER","tag_article_rel","tag_article_rel").Find(&articles)
	err=engine.Where("Ctid =?",ctid).Join("INNER","tag_article_rel","tag_article_rel.cid=article.id").Desc("id").Find(&articles)

	return
}

// GetArticleByTag 通过标签列表
func (article *Article) PageArticleByTag(ctid ,offset, limit int64) (articles []*Article, err error) {
	// err=engine.Join("INNER","tag_article_rel","tag_article_rel").Find(&articles)
	err=engine.Where("Ctid =?",ctid).Join("INNER","tag_article_rel","tag_article_rel.cid=article.id").Desc("id").Limit(int(limit), int(offset)).Find(&articles)
	return
}

// Count 统计所有文章的数量
func (article *Article) Count() (count int64,err error) {
	count,err= engine.Count(article)
	return
}

