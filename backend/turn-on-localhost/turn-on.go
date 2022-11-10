package main

/*
It's a test go source file in order to turn on localhost network


*/
import (
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)



func main() {
	e := echo.New()
	e.HideBanner = true
	E := e.Group("x")
	E.GET("/hello", func (c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})

	r := e.Group("y")
	r.GET("/Ping", func (c echo.Context) error {
		return c.String(http.StatusOK, "Pong!")
	})

	e.GET("/hello", func (c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})
	
	e.GET("/Ping", func (c echo.Context) error {
		return c.String(http.StatusOK, "Pong!")
	})


	e.Start("127.0.0.1:80")
	fmt.Println("End!")

}
/*

cmd运行go run turn-on.go后即启动服务器
重新开一个cmd
运行curl localhost/Ping -v
即连接到服务器，并通过Ping测试
总路由开启前可以多次调用GET，获取的方法会同时生效


*/