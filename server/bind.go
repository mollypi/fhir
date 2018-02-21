package server

import (
	"github.com/labstack/echo"
)

func FHIRBind(c echo.Context, i interface{}) (err error) {
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}

	return c.Bind(i)
}
