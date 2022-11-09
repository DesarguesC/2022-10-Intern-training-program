package todo

import (
	"2022-10-Intern-training-program/backend/real-backend/app/response"
	"github.com/labstack/echo/v4"
)

var cnt int64

func add(c echo.Context, tit string, con string) {
	if cnt == 0 {
		cnt = 1
	}
	x := table.Todos{cnt, 111, tit, con}
	// id ?
	cnt++
	er := db.Debug().Create(&x).Error
	if er != nil {
		// logrus.Panic("Cannot create object Todos.")
		response.ToResponse(c, 200, "Cannot create object Todos.", tit, con)
	}

}
