package dbcontroller

// 启动服务器，链接，并创建Users和Todos两个数据库

import (
	"Fly/module/dbcontroller/base"
	"github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB_Init() {
	uri := ConnectMake()
	//
	base.Db, base.Err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	if base.Err != nil {
		logrus.Fatal(base.Err)
	}
	if base.Err != nil {
		logrus.Panic("strange!")
	}

	base.Err = base.Db.AutoMigrate(&base.Users{})
	if base.Err != nil {
		// logrus.Fatal(a_err)
		logrus.Panic("UNABLE to create Users!")
	}
	base.Err = base.Db.AutoMigrate(&base.Todos{})
	if base.Err != nil {
		logrus.Panic("UNABLE to create Todos!")
	}
}
