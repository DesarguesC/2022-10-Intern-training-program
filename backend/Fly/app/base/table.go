package base

import "gorm.io/gorm"

// table

var (
	Db  *gorm.DB
	Err error
)

/*
这两个变量必须全局，使用":="声明会被认为是临时变量
之后遭到释放将出现空指针错误
*/
// ↑ utf8
/*
Both "Db" and "Err" must be declared globally or they'll be
released. It leads to NULL pointer exception
*/

type Users struct {
	//gorm.Model
	Id     uint `gorm:"primary_key"`
	Name   string
	Passwd string
}

type Todos struct {
	Id      uint `gorm:"promary_key"`
	User_id uint
	Tittle  string
	Content string
}
