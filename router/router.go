package router

import (
	"net/http"

	"github.com/labstack/echo"
)

func Routing(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
