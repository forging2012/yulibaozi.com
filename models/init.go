package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var (
	engine  *xorm.Engine
	ArticleColumn []string
)

func init() {
	var err error
	dateSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", "root", "mypassword", "39.108.50.89", 3306, "yulibaozi") + "&loc=Asia%2FShanghai"
	engine, err = xorm.NewEngine("mysql", dateSource)
	if err != nil {
		fmt.Println("初始化数据库连接失败，err:", err)
		return
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(5)  //设置连接池的空闲数大小
	engine.SetMaxOpenConns(30) //设置最大打开连接数
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "tbl_")
	engine.SetTableMapper(tbMapper)
	engine.CreateTables(
		new(Article),
		new(Category),
		new(EmailInfo),
		new(SmtpConfig),
		new(User),
		new(WebConfig),
	)
	ArticleColumn = engine.TableInfo(new(Article)).ColumnsSeq()
	for _, v := range ArticleColumn{
		fmt.Println(v)
	}
}

func GetEngine() *xorm.Engine {
	return engine
}
