package net_utils

import (
	"github.com/labstack/echo/v4"
)

func QueryParam(c echo.Context) error {
	query := c.QueryParam("query")
	// QueryParam-查询参数		Param-路径参数
	// (↑utf-8)
	//fmt.Println(query)
	return c.String(http.StatusOK, query)
}

// Print parameter 'query' directly