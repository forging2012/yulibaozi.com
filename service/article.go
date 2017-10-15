package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yulibaozi/yulibaozi.com/cache"
	"github.com/yulibaozi/yulibaozi.com/conf"
	initial "github.com/yulibaozi/yulibaozi.com/initialization"
	"github.com/yulibaozi/yulibaozi.com/models"
)

type ArticleService struct{}

func (articleService *ArticleService) Update(article *models.Article) (updId int64, err error) {
	updId, err = article.Update()
	if err != nil {
		return
	}
	//更新redis
	if updId > 0 {
		err = new(cache.RedisUtils).HMSET(fmt.Sprintf(conf.ArticleFormat, article.Id), article)
		return
	}
	err = fmt.Errorf("插入数据失败")
	return
}

// Inset 写入文章
//写入文章的时候 先先写入数据库在写入缓存，写入缓存的时候同时写入hash和zset
func (articleService *ArticleService) Inset(article *models.Article) (newId int64, err error) {
	newId, err = article.Inset()
	if err != nil {
		return
	}
	if newId <= 0 {
		err = fmt.Errorf("插入数据失败")
		return
	}
	//写入缓存
	key := fmt.Sprintf(conf.ArticleFormat, newId)
	err = new(cache.RedisUtils).HMSETWEIGHT(key, conf.ArticleZset, newId, article)
	return
}

// Delete 删除文章并删除redis里面
func (articleService *ArticleService) Delete(article *models.Article) (delId int64, err error) {
	delId, err = article.Delete()
	if err != nil {
		return
	}
	err = new(cache.RedisUtils).HDELWEIGHT(fmt.Sprintf(conf.ArticleFormat, article.Id), conf.ArticleZset, initial.ArticleColumn)
	return
}

// GetArticle 获取某一条文章
func (articleService *ArticleService) GetArticle(aid int64) (article *models.Article, err error) {
	var (
		ok bool
	)
	err = new(cache.RedisUtils).HGETALL(fmt.Sprintf(conf.ArticleFormat, aid), article)
	if err != nil {
		ok, err = article.GetOne(aid)
		if !ok {
			err = fmt.Errorf("没有找到数据数据")
			return
		}
	}
	return
}

// PageArticle 分页获取article列表
func (articleService *ArticleService) PageArticle(offset int, limit int) (articles []interface{}, err error) {
	var (
		ok bool
	)
	redisUtils := &cache.RedisUtils{}
	limits := offset + limit - 1
	var (
		article       *models.Article
		articleStruct []*models.Article
	)
	slice, err := redisUtils.GetZsetFields(conf.ArticleZset, offset, limits)
	if err != nil {
		goto ERRDISPOSE
	}
	for _, v := range slice {
		//如果在redis没有查找到数据，就去数据库补
		err = redisUtils.HGETALL(v, &article)
		if err != nil {
			articleIndex := strings.Split(v, "article:")
			if len(articleIndex) != 2 {
				continue
			} else {
				aid, _ := strconv.ParseInt(articleIndex[1], 10, 64)
				ok, err = article.GetOne(aid)
				if ok == false || err != nil {
					continue
				}
				articles = append(articles, article)
				continue
			}
		}
		articles = append(articles, article)
	}
	if err == nil && len(articles) >= 0 {
		return
	}
	//查数据库
ERRDISPOSE:
	articleStruct, err = article.PageArticle(offset, limit)
	if err != nil {
		return
	} else {
		for _, v := range articleStruct {
			articles = append(articles, &v)
		}
	}

	return
}

// PageArticleByCid 根据filter获取文章列表，根据时间倒序排列
// filter 是分类id
func (articleService *ArticleService) PageArticleByCid(offset int, limit int, filter int64) (articles []interface{}, err error) {
	return
}

func (articleService *ArticleService) ArticleLike(cid int64) (articles []interface{}, err error) {
	return
}

// ArticlesByTag 通过标签编号获取文章列表
// func (articleService *ArticleService) ArticlesByTag(tid int64) (articles []*models.Article, err error) {
// 	//先查询redis
// 	key := fmt.Sprintf(conf.CatetagtoarticleFormat, tid)
// 	article := &models.Article{}
// 	redisUtils := &cache.RedisUtils{}
// 	slice, err := redisUtils.Keys(key)
// 	if len(slice) <= 0 || err != nil {
// 		//查询数据库
// 		articles, err = new(models.Article).GetArticleByTag(tid)
// 		return
// 	} else {
// 		for _, v := range slice {
// 			err = redisUtils.HGETALL(v, article)
// 			if err != nil {
// 				//如果redis没有查出来数据，就去数据库补
// 				idStrs := strings.SplitAfterN(v, ":", 3)
// 				if len(idStrs) != 2 {
// 					fmt.Println("分离错误!")
// 					continue
// 				} else {
// 					id, err := strconv.ParseInt(idStrs[1], 10, 64)
// 					if err != nil {
// 						fmt.Println("id转换错误", err.Error())
// 						continue
// 					} else {
// 						ok, err := new(models.Article).GetOne(id)
// 						if ok == false || err != nil {
// 							fmt.Println("获取数据库错误，", err.Error())
// 							continue
// 						} else {
// 							articles = append(articles, article)
// 						}
// 					}
// 				}
// 				continue
// 			} else {
// 				articles = append(articles, article)
// 			}

// 		}
// 	}

// 	return
// }

// // Count 获取文章总数
// func (articleService *ArticleService) Count() (count int64, err error) {
// 	count, err = new(cache.RedisUtils).CountByKey(conf.ArticleZset)
// 	if err != nil {
// 		//查询数据库
// 		// count, err = new(models.Article).Count()
// 		// return
// 	}
// 	return
// }

// //根据tagId 查询文章总数
// // func (articleService *ArticleService) CountByTagId(ctid int64) (count int64, err error) {
// // 	slice := []string{}
// // 	slice, err = new(cache.RedisUtils).Keys(fmt.Sprintf(conf.CatetagtoarticleFormat, ctid))
// // 	if err != nil {
// // 		//查数据库
// // 		count, err = new(models.CategoryAndTag).GetCount(ctid)
// // 		return
// // 	}
// // 	count = int64(len(slice))
// // 	if count == 0 {
// // 		//查数据库
// // 		count, err = new(models.CategoryAndTag).GetCount(ctid)
// // 		return
// // 	}
// // 	return
// // }
// func (articleService *ArticleService) PageArticleByTid(tid, offset, limit int64) (articles []*models.Article, err error) {
// 	// 先redis
// 	redisUtils := &cache.RedisUtils{}
// 	limits := offset + limit - 1
// 	var (
// 		article *models.Article
// 	)
// 	slice, err := redisUtils.GetZsetREVFields(conf.ArticleZset, offset, limits)
// 	if err != nil {
// 		goto ERRDISPOSE
// 	} else {
// 		for _, v := range slice {
// 			err = redisUtils.HGETALL(v, article)
// 			if err != nil {
// 				//如果redis没有查出来数据，就去数据库补
// 				idStrs := strings.SplitAfterN(v, ":", 3)
// 				if len(idStrs) != 2 {
// 					fmt.Println("分离错误!")
// 					continue
// 				} else {
// 					id, err := strconv.ParseInt(idStrs[1], 10, 64)
// 					if err != nil {
// 						fmt.Println("id转换错误", err.Error())
// 						continue
// 					} else {
// 						ok, err := new(models.Article).GetOne(id)
// 						if ok == false || err != nil {
// 							fmt.Println("获取数据库错误，", err.Error())
// 							continue
// 						} else {
// 							articles = append(articles, article)
// 						}
// 					}
// 				}
// 				continue
// 			} else {
// 				articles = append(articles, article)
// 			}

// 		}
// 	}
// 	//查数据库 并写入数据库
// ERRDISPOSE:
// 	articles, err = article.PageArticleByTag(tid, offset, limit)
// 	for _, v := range articles {
// 		redisUtils.ZADD(conf.ArticleZset, fmt.Sprintf(conf.WriteArticleAndTag, tid, v.Id), v.Id)
// 		redisUtils.HMSET(fmt.Sprintf(conf.ArticleFormat, v.Id), v)
// 	}
// 	return

// }

func (articleService *ArticleService) GetOne(aid int64) (article *models.Article, err error) {
	var ok bool
	ok, err = article.GetOne(aid)
	if err != nil {
		return
	}
	if !ok {
		err = fmt.Errorf("查询文章:%d,出错。", aid)
	}
	return
}

// Total 获取文章总数
func (articleService *ArticleService) Total() (count int64, err error) {
	count, err = new(models.Article).Total()
	return
}

// Hot 更加文章浏览数排出最热文章
func (articleService *ArticleService) Hot() (list []interface{}, err error) {
	return
}

func (articleService *ArticleService) UpdateViewCount(article *models.Article) (err error) {
	err = article.UpdateViewCount()
	return
}
