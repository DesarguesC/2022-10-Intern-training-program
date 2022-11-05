package todo

import(
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"model/table"
	"app/response"
)
var cnt_ int64

func get(c echo.Context, idd int) {
	x := table.Todos{}
	x.Id = idd
	er := db.Debug().Find(&x).Error
	if er != nil {
		// logrus.Panic("Find no id = " + string(idd))
		response.ToResponse(c, 200, "Find no id = "+ + string(idd),)
	}
}
