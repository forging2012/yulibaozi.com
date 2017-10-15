package main

import (
	"fmt"

	_ "github.com/yulibaozi/yulibaozi.com/initialization"
	"github.com/yulibaozi/yulibaozi.com/routers"

	"github.com/devfeel/dotweb"
)

func main() {
	fmt.Println("======yulibaozi.com======")
	app := dotweb.New()
	app.SetDevelopmentMode()
	routers.InitRoute(app.HttpServer)
	app.SetEnabledLog(true)
	app.SetLogPath("logs")
	port := 8080
	err := app.StartServer(port)
	if err != nil {
		fmt.Println("启动服务失败...")
	}
	fmt.Println("服务已经启动....")
}
