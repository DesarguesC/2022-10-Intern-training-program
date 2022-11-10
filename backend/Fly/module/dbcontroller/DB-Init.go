package dbcontroller

// 启动服务器，链接，并创建Users和Todos两个数据库

import (
	"Fly/module/dbcontroller/base"
	"github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
