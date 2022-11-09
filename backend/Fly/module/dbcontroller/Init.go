package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)
/*
这两个变量必须全局，使用":="声明会被认为是临时变量
之后遭到释放将出现空指针错误
*/
// ↑ utf8
/*
Both "db" and "err" must be declared globally or they'll be 
released. It leads to NULL pointer exception
*/

func Init() {
	uri := ConnectMake()
	db, err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	if err != nil {
		logrus.Panic("strange!")
	}
	a_err := db.AutoMigrate(&Users{})
	if a_err != nil {
		// logrus.Fatal(a_err)
		logrus.Panic("UNABLE to create Users!")
	}
	b_err := db.AutoMigrate(&Todos{})
	if b_err != nil {
		logrus.Panic("UNABLE to create Todos!")
	}
}

func ConnectMake() string {
	viper.SetConfigName("uri")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	// which can be called multiple times to add possible search options
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
	loginInfo := viper.GetStringMapString("param")

	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["database"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	// NEVER use uppercase --- QwQ
	// "@(localhost)/" or "@(127.0.0.1:3306)/"  ---> OK
	return dbArgs
}
