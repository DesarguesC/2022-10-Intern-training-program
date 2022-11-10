package dbcontroller

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 与数据库建立连接

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
