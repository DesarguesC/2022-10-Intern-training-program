package controller

import(
	"github.com/labstack/echo/v4";
	"net/http";
	"fmt"
)

func Ping(c echo.Context) error {
	fmt.Println("pong!")
	return c.String(http.StatusOK, "pong!")
}
// Ping: for judgement
// 不加Println使用cmd访问也可以直接输出
// utf8