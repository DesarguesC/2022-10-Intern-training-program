package base

import (
	"log"
)

func Todos_insert(id uint, uid uint, ti string, con string) {
	a_todo := Todos{id, uid, ti, con}
	Err = Db.Debug().Create(&a_todo).Error
	if Err != nil {
		log.Fatal(Err)
	}
}

func Todos_delete(id uint) {
	a_todo := Todos{Id: id}
	Err = Db.Debug().Delete(&a_todo).Error
	if Err != nil {
		log.Fatal(Err)
	}
}

func Todos_query(uid uint) Todos {
	a_todo := Todos{}
	a_todo.User_id = uid
	Err = Db.Debug().Find(&a_todo).Error
	if Err != nil {
		log.Fatal(Err)
	}
	// return ???
	return a_todo
}

func Todos_update(id uint, con string) {
	a_todo := Todos{}
	a_todo.Id = id
	a_todo.Content = con
	Err = Db.Debug().Updates(&a_todo).Error
	if Err != nil {
		log.Fatal(Err)
	}
}
