package main

import (
	"fmt";
	"log";
	"gorm.io/driver/mysql";
	"gorm.io/gorm";
  )
type Users struct {
	//gorm.Model
	Id uint `gorm:"primary_key"`
	Name string
	Passwd string
}
type Todos struct{
	Id uint `gorm:"promary_key"`
	User_id uint
	tittle string
	content string
}
// create table
  
func main() {
	url := "root:Cd2003527!?#@(127.0.0.1:3306)/todo?charset=utf8&parseTime=True&loc=Local"
	// database name: todo
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		// or another: panic("....(error messages)")
	}

	db.AutoMigrate(&Users{})
	// 	create table: Users

	// insertion
	a_user := Users{1,"chen","12321"}
	err = db.Debug().Create(&a_user).Error
	if err != nil {
		log.Fatal(err)
	}
	b_user := Users{3,"ding","xxx"}
	c_user := Users{4,"qwq","233333"}
	err = db.Debug().Create(&b_user).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Create(&c_user).Error
	if err != nil {
		log.Fatal(err)
	}


	// query
	x_user := Users{}
	x_user.Id = 3
	err = db.Debug().Find(&x_user).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x_user)

	// update
	y_user := Users{}
	y_user.Id = 4
	y_user.Passwd = ("333332")	// StringFrom ?
	err = db.Debug().Updates(&y_user).Error
	if err != nil {
		log.Fatal(err)
	}

	// delete
	z_user := Users{Id:3}
	err = db.Debug().Delete(&z_user).Error
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Todos{})
	// create table: Todos
	// the same

	// insertion
	aa_user := Todos{1,123,"chen","12321"}
	err = db.Debug().Create(&aa_user).Error
	if err != nil {
		log.Fatal(err)
	}
	bb_user := Todos{3,111,"ding","xxx"}
	cc_user := Todos{4,233,"ok","233333"}
	err = db.Debug().Create(&bb_user).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Create(&cc_user).Error
	if err != nil {
		log.Fatal(err)
	}

	// query
	xx_user := Todos{}
	xx_user.Id = 3
	err = db.Debug().Find(&xx_user).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x_user)

	// update
	yy_user := Todos{}
	yy_user.Id = 4
	yy_user.content = ("333332")	// StringFrom ?
	err = db.Debug().Updates(&yy_user).Error
	if err != nil {
		log.Fatal(err)
	}

	// delete
	zz_user := Todos{Id:3}
	err = db.Debug().Delete(&zz_user).Error
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("All've been finished.")

}

