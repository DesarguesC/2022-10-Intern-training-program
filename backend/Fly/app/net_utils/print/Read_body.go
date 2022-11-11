package print

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	//"github.com/go-playground/validator/v10"
	"log"
)

func Read_body(c echo.Context) error {
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&body)
	x := string(body)
	return c.String(http.StatusOK, x)
	// 直接打印
}
