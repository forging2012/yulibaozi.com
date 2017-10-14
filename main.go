package main

import (
	"fmt"
	_ "github.com/yulibaozi/yulibaozi.com/models"
	"github.com/yulibaozi/yulibaozi.com/routers"
	

	"github.com/devfeel/dotweb"
)

func main() {
	
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
}
