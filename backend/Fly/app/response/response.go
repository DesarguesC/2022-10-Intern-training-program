package response

// response 格式化返回错误信息
// 用来处理 err != nil 时的信息

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MResponse(c echo.Context, code int64, msg string, data ...interface{}) error {
	return c.JSON(http.StatusOK, Err_msg{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

//	Replace original c.String
