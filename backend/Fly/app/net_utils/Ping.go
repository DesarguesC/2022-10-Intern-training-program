package net_utils
// 请求处理模块
import (
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}
// Ping: for judgement

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Helo World!")
}
// for entertainment