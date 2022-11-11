package middleware

import (
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("uid", 114514)
		return next(c) // 直接通过当前用户请求
	}
}

// BONUS

func Check_Info(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}
