package routers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/controllers"
)

// init 初始化路由
func InitRoute(server *dotweb.HttpServer)  {
	// server.Router().GET("/home",)
	 userController:=&controllers.UserController{}
	 fileController:=&controllers.FileController{}
	
	server.Router().POST("/post",userController.InsetOrUpdate)
	server.Router().GET("/del",userController.Delete)
	server.Router().GET("/get",userController.GetOne)

	server.Router().GET("/fileget",fileController.Show)
	server.Router().POST("/filepost",fileController.Insert)
}
