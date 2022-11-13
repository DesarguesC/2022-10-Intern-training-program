package middleware

import (
	"Fly/app/response"
	"Fly/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("uid", 114514)
		return next(c) // 直接通过当前用户请求
	}
}

//	middleware: all users are authorized

// BONUS

func Check_Info(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid, err := strconv.Atoi(c.Param("uid"))
		passwd := c.Param("passwd")
		if err != nil {
			log.Panic("Invalid User-Id!")
		}

		one := &utils.PermitPerson{uid, passwd}
		valid := validator.New()
		fatal_err := valid.Struct(one)
		if fatal_err != nil {
			viper.SetConfigName("loginErr")
			viper.SetConfigType("json")
			viper.AddConfigPath("Fly/app/middleware/config/")
			if errrr := viper.ReadInConfig(); errrr != nil {
				log.Fatal(errrr)
			}
			Msg := viper.GetStringMapString("loginErr")
			var (
				er_ error
				co  int
			)
			co, er_ = strconv.Atoi(Msg["code"])
			if er_ != nil {
				log.Fatal(er_)
			} // 该错误应该是不会发生的
			response.MResponse(c, int64(co), Msg["msg"])
		}
		return response.MResponse(c, http.StatusOK, "Welcome: User-"+string(uid))
	}
}

//  Only users with valid id and password are authorized
