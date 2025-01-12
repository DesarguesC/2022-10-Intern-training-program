package main

import (
	"fmt";
	"log";
  )
func main() {

	Init()
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

