package conf

var (
	ArticleFormat          = "hasharticle:%d"        //redis 文章key
	CatetagFormat          = "catetag:%d"            // redis 分类key
	ArticletocatetagFormat = "catetag:*:article:%d"  //redis 通过文章查所有的分类和标签
	CatetagtoarticleFormat = "catetag:%d:article:*"  //redis 通过分类查找文章
	WriteArticleAndTag     = "catetag:%d:article:%d" //写入redis的key
	ArticlMail             = "article:%d:mail:*"     //文章评论写入的key
	ArticleZset            = "zsetarticle"           //文章zet 用来分页

)


