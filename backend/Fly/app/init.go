package app

import (
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

	// entertain
	E.GET("/Ping", net_utils.Ping)
	E.GET("/Hello", net_utils.Hello)

	// 接受参数并直接打印
	E.GET("/:param", print.QueryP)

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
	E.GET("/api/login/:uid/:passwd", net_utils.Check_Login)

	E.Start("127.0.0.1:80") // localhost端口

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
	E.Logger.Fatal(E.Start(":1323"))
}
