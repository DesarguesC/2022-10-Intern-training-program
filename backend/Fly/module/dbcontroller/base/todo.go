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

func Todos_query(id uint) {
	a_todo := Todos{}
	a_todo.Id = id
	Err = Db.Debug().Delete(&a_todo).Error
	if Err != nil {
		log.Fatal(Err)
	}
}

func Todos_update(id uint, con string) {
	a_todo := Todos{}
	a_todo.Id = id
	a_todo.content = con
	Err = Db.Debug().Updates(&a_todo).Error
	if Err != nil {
		log.Fatal(Err)
	}
}
