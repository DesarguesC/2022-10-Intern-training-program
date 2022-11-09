package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func InitWebWork() {
	e = echo.New()
	e.GET("/hello", read_body)
	e.Start("127.0.0.1:80")
	fmt.Println("first")
}
