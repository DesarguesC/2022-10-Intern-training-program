package net_utils

import (
	"Fly/module/dbcontroller/base"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

var users_cnt uint
var todos_cnt uint

func Add_todo(c echo.Context) error {
	ti := c.QueryParam("ti")
	con := c.QueryParam("con")

	base.Todos_insert(users_cnt, todos_cnt, ti, con)

	users_cnt++
	todos_cnt++

	return c.String(http.StatusOK, "added: Todos Id = "+string(todos_cnt-1))
	// 好像不用fmt也可以直接打印

	// QueryParam-查询参数		Param-路径参数
	// (↑ utf-8)
}

func Get_todo(c echo.Context) error {
	uid := c.Param("uid")
	uu, er := strconv.Atoi(uid)
	if er != nil {
		log.Panic("Invalid User_Id")
	}
	x := base.Todos_query(uint(uu))
	return c.String(http.StatusOK, "Todos: [id = "+string(x.Id)+"], [tittle = "+x.Tittle+"], [content = "+x.Content+"]")

}
