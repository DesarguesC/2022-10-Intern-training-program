package net_utils

import (
	"Fly/app"
	"Fly/app/middleware"
)

func Check() {
	api := app.E.Group("api")
	// 一个组路由
	api.Use(middleware.Auth)
	api.GET("/Ping", Ping)
	// 组路由的运行并不会持续，这里的Ping会在
}
