package print

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func QueryP(c echo.Context) error {
	query := c.QueryParam("param")
	return c.String(http.StatusOK, query)
	// 好像不用fmt也可以直接打印

	// QueryParam-查询参数		Param-路径参数
	// (↑ utf-8)
}

// Print parameter 'query' directly
