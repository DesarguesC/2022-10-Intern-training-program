package app

import (
	"github.com/labstack/echo/v4"
	"Fly/app/net_utils"
	"net/http"
	//"github.com/go-playground/validator";
	//"encoding/json";
	"log"
	//"flybitch/go-net/model"
)

var e *echo.Echo

// 创建echo实例 -> 注册中间件 -> 注册用户路由 -> 启动http server -> accept请求

func InitNetWork() {
	e = echo.New()
	e.HideBanner = true // 隐藏startup message

	e.Validator = 
	

}

func StartServer() {
	e.Logger.Fatal(e.Start("..."))
}



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





// Print parameter 'query' directly

// func read_body(rw http.ResponseWriter, req *http.Request) {
// 	length := req.ContentLength
// 	body := make([]byte, length)
// 	req.Body.Read(body)
// 	fmt.Fprintln(rw, body)
// 	fmt.Println(body)
// }


