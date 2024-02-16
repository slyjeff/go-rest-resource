package main

import (
	"github.com/labstack/echo/v4"
	"github.com/slyjeff/rest-resource/encoding"
	"github.com/slyjeff/rest-resource/resource"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var r resource.Resource
		r.Data("message", "Hello World!")

		return respond(c, r)
	})
	e.Logger.Fatal(e.Start(":80"))
}

func respond(c echo.Context, r resource.Resource) error {
	value, contentType := encoding.ForAcceptHeader(r, c.Request().Header)
	c.Response().Header().Set("Content-Type", contentType)
	return c.String(http.StatusOK, value)
}
