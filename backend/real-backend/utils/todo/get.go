package todo

import (
	"app/response"
	"github.com/labstack/echo/v4"
	"model/table"
)

var cnt_ int64

func get(c echo.Context, idd int) {
	x := table.Todos{}
	x.Id = idd
	er := db.Debug().Find(&x).Error
	if er != nil {
		// logrus.Panic("Find no id = " + string(idd))
		response.ToResponse(c, 200, "Find no id = "+string(idd))
	}
}
