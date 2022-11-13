package main

import (
	"Fly/app"
	"Fly/module/dbcontroller"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {

	fmt.Println("Hello World!")
	// Entertain
	logrus.SetReportCaller(true)
	dbcontroller.DB_Init()
	app.InitNetWork()
	app.StartServer()
}
