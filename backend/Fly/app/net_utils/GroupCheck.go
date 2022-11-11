package net_utils

import (
	"Fly/app"
	"Fly/app/middleware"
	"Fly/app/net_utils/print"
)

func Check() {
	api := app.E.Group("api")
	// 一个组路由
	api.Use(middleware.Auth)
	api.GET("/Ping", Ping)
	api.GET("/", print.Read_body)
	// 组路由调用并不会持续，这里的handle调用一次后就会失效继续执行之后的进程
}
