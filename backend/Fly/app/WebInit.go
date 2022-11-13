package app

import (
	"Fly/app/middleware"
	"Fly/app/net_utils"
	"Fly/app/net_utils/print"
	"Fly/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var E *echo.Echo // 供全局使用的echo变量

/*
创建echo实例 -> 注册中间件 -> 注册用户路由 -> 启动http server -> accept请求
*/

func InitNetWork() {
	E = echo.New()
	E.HideBanner = true // 隐藏startup message
	net_utils.Check()
	// test

	E.GET("/Ping", net_utils.Ping)
	E.GET("/Hello", net_utils.Hello)
	// entertain

	E.GET("/:param", print.QueryP) // 接受参数并直接打印

	//	complex api ↓	can be seen in Fly/app/net_utils/API.go
	E.GET("/api/todo/add/:ti/:con", net_utils.Add_todo)
	/*
		e.g.	localhost/api/todo/add/:using/:hhhh
		return -> "added: Todos Id = ..."
	*/
	E.GET("/api/todo/get/:uid", net_utils.Get_todo)
	/*
		e.g.	localhost/api/todo/get/:233
		return -> "Todos: [id=233][...][...]..."
	*/
	E.Use(middleware.Check_Info)
	E.Validator = &utils.CustomValidator{}
	logrus.Info("echo framework initialized")
}

/*

cmd运行go run turn-on.go后即启动服务器
重新开一个cmd
运行curl localhost/Ping -v
即连接到服务器，并通过Ping测试
总路由开启前可以多次调用GET，获取的方法会同时生效


*/

func StartServer() {
	E.Logger.Fatal(E.Start("127.0.0.1:80"))
	// start localhost
}
