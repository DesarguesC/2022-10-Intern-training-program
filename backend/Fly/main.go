package main

import (
	"Fly/module/dbcontroller"
	"fmt"
)

func main() {
	dbcontroller.DB_Init()
	fmt.Println("Hello World!")
}
