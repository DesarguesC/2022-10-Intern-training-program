package go_net

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	//"github.com/go-playground/validator";
	//"encoding/json";
	"log"
	//"flybitch/go-net/model"
)

//func main() {

//e := echo.New()
//e.GET("/hello", read_body)
//e.Start("127.0.0.1:80")
//fmt.Println("first")
//
// e.GET("/xyz", QueryParam)
// e.Start("127.0.0.1:80")
// fmt.Println("second")

// server := http.Server {
// 	Addr: "127.0.0.1:80",
// }
// http.HandleFunc("/hello-s", read_body)
// server.ListenAndServe()
// fmt.Println("third")

//}

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}

// Ping: for judgement
func QueryParam(c echo.Context) error {
	query := c.QueryParam("query")
	// QueryParam-查询参数		Param-路径参数
	// (↑utf-8)
	//fmt.Println(query)
	return c.String(http.StatusOK, query)
}

// Print parameter 'query' directly

// func read_body(rw http.ResponseWriter, req *http.Request) {
// 	length := req.ContentLength
// 	body := make([]byte, length)
// 	req.Body.Read(body)
// 	fmt.Fprintln(rw, body)
// 	fmt.Println(body)
// }

func read_body(c echo.Context) error {
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&body)
	x := string(body)
	return c.String(http.StatusOK, x)
}
