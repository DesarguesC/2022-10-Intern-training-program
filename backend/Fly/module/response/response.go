package response

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
