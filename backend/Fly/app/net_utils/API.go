package net_utils

import (
	"Fly/app/base"
	"Fly/app/response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

var users_cnt uint
var todos_cnt uint

func Add_todo(c echo.Context) error {
	ti := c.QueryParam("ti")
	con := c.QueryParam("con")

	error_ := base.Todos_insert(users_cnt, todos_cnt, ti, con)

	if error_ != nil {
		viper.SetConfigName("insertErrmsg")
		viper.SetConfigType("json")
		viper.AddConfigPath("Fly/module/dbcontroller/config/todoErr/")
		if errrr := viper.ReadInConfig(); errrr != nil {
			log.Fatal(errrr)
		}
		Msg := viper.GetStringMapString("insertErrmsg")
		var (
			er_ error
			co  int
		)
		co, er_ = strconv.Atoi(Msg["code"])
		if er_ != nil {
			log.Fatal(er_)
		} // 该错误不可能发生
		response.MResponse(c, int64(co), Msg["msg"])
	}

	users_cnt++
	todos_cnt++
	return response.MResponse(c, http.StatusOK, "added: Todos Id = "+string(todos_cnt-1))
	//return c.String(http.StatusOK, "added: Todos Id = "+string(todos_cnt-1))
	// 好像不用fmt也可以直接打印

	// QueryParam-查询参数		Param-路径参数
	// (↑ utf-8)
}

func Get_todo(c echo.Context) error {
	uid := c.Param("uid")
	uu, er := strconv.Atoi(uid)
	if er != nil {
		log.Panic("Invalid User_Id: User_Id value must be an Integer")
	}
	x, er_ := base.Todos_query(uint(uu))
	if er_ != nil {
		viper.SetConfigName("queryErrmsg")
		viper.SetConfigType("json")
		viper.AddConfigPath("Fly/module/dbcontroller/config/todoErr/")
		if errrr := viper.ReadInConfig(); errrr != nil {
			log.Fatal(errrr)
		}
		Msg := viper.GetStringMapString("queryErrmsg")
		var (
			er_ error
			co  int
		)
		co, er_ = strconv.Atoi(Msg["code"])
		if er_ != nil {
			log.Fatal(er_)
		} // 该错误不可能发生
		response.MResponse(c, int64(co), Msg["msg"])
	}
	return response.MResponse(c, http.StatusOK, "Todos: [id = "+string(x.Id)+"], [tittle = "+x.Tittle+"], [content = "+x.Content+"]")
	//return c.String(http.StatusOK, "Todos: [id = "+string(x.Id)+"], [tittle = "+x.Tittle+"], [content = "+x.Content+"]")

}

//
//func Check_Login(c echo.Context) error {
//	uid, err := strconv.Atoi(c.Param("uid"))
//	passwd := c.Param("passwd")
//	if err != nil {
//		log.Panic("Invalid User-Id!")
//	}
//
//	one := &utils.PermitPerson{uid, passwd}
//	valid := validator.New()
//	fatal_err := valid.Struct(one)
//	if fatal_err != nil {
//		viper.SetConfigName("loginErr")
//		viper.SetConfigType("json")
//		viper.AddConfigPath("Fly/app/middleware/config/")
//		if errrr := viper.ReadInConfig(); errrr != nil {
//			log.Fatal(errrr)
//		}
//		Msg := viper.GetStringMapString("loginErr")
//		var (
//			er_ error
//			co  int
//		)
//		co, er_ = strconv.Atoi(Msg["code"])
//		if er_ != nil {
//			log.Fatal(er_)
//		} // 该错误应该是不会发生的
//		response.MResponse(c, int64(co), Msg["msg"])
//	}
//
//	return response.MResponse(c, http.StatusOK, "Welcome: User-"+string(uid))
//
//}
