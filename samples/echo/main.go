package main

import (
	"net/http"

	"github.com/dimorinny/gopermission"
	"github.com/labstack/echo"
)

type HasQwertyHeader struct{}

func (ch HasQwertyHeader) HasPermission(request *http.Request) bool {
	return request.Header.Get("Qwerty") != ""
}

func handler(c *echo.Context) error {
	c.String(200, "Granted")
	return nil
}

func main() {
	permission := gopermission.New(HasQwertyHeader{})
	qwertyMiddleware := func(c *echo.Context) error {
		if !permission.IsPermitted(c.Request()) {
			c.String(403, "Error")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return nil
	}

	echo := echo.New()
	echo.Use(qwertyMiddleware)
	echo.Get("/", handler)
	echo.Run(":9090")
}
