package print

import (
	"fmt";
	"net/http";
	"github.com/labstack/echo/v4";
	"io/ioutil";
	"log"
)

func read_body(c echo.Context) error {
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	if  err != nil {
		log.Fatal(err)
	}
	fmt.Println(&body)
	x := string(body)
	return c.String(http.StatusOK, x)
}