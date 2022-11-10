package net_utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func QueryParam(c echo.Context) error {
	query := c.QueryParam("query")
	return c.String(http.StatusOK, query)
	// QueryParam-查询参数		Param-路径参数
	// (↑ utf-8)
}

// Print parameter 'query' directly
