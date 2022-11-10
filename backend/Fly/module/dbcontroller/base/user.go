package base

import (
	"log"
)

func Users_insert(id uint, name string, passwd string) {
	a_user := Users{id, name, passwd}
	Err = Db.Debug().Create(&a_user).Error
	if Err != nil {
		log.Fatal(Err)
	}
}

func Users_delete(id uint) {
	a_user := Users{Id: id}
	Err = Db.Debug().Delete(&a_user).Error
	if Err != nil {
		log.Fatal(Err)
	}
}

// search by Id
func Users_query(id uint) {
	a_user := Users{}
	a_user.Id = id
	Err = Db.Debug().Find(&a_user).Error
	if Err != nil {
		log.Fatal(Err)
	}
}

// update after searching by Id, update passwd
func Users_update_passwd(id uint, passwd string) {
	a_user := Users{}
	a_user.Id = id
	a_user.Passwd = passwd
	Err = Db.Debug().Updates(&a_user).Error
	if Err != nil {
		log.Fatal(Err)
	}
}

// func main() {

// 	Init()
// 	// insertion

// 	a_user := Users{1,"chen","12321"}
// 	Err = Db.Debug().Create(&a_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}
// 	b_user := Users{3,"ding","xxx"}
// 	c_user := Users{4,"qwq","233333"}
// 	Err = Db.Debug().Create(&b_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}
// 	Err = Db.Debug().Create(&c_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}

// 	// query
// 	x_user := Users{}
// 	x_user.Id = 3
// 	Err = Db.Debug().Find(&x_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}
// 	fmt.Println(x_user)

// 	// update
// 	y_user := Users{}
// 	y_user.Id = 4
// 	y_user.Passwd = ("333332")	// StringFrom ?
// 	Err = Db.Debug().Updates(&y_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}

// 	// delete
// 	z_user := Users{Id:3}
// 	Err = Db.Debug().Delete(&z_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}

// 	// the same
// 	// insertion
// 	aa_user := Todos{1,123,"chen","12321"}
// 	Err = Db.Debug().Create(&aa_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}
// 	bb_user := Todos{3,111,"ding","xxx"}
// 	cc_user := Todos{4,233,"ok","233333"}
// 	Err = Db.Debug().Create(&bb_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}
// 	Err = Db.Debug().Create(&cc_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}

// 	// query
// 	xx_user := Todos{}
// 	xx_user.Id = 3
// 	Err = Db.Debug().Find(&xx_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}
// 	fmt.Println(x_user)

// 	// update
// 	yy_user := Todos{}
// 	yy_user.Id = 4
// 	yy_user.content = ("333332")	// StringFrom ?
// 	Err = Db.Debug().Updates(&yy_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}

// 	// delete
// 	zz_user := Todos{Id:3}
// 	Err = Db.Debug().Delete(&zz_user).Error
// 	if Err != nil {
// 		log.Fatal(Err)
// 	}

// 	fmt.Println("All've been finished.")

// }
