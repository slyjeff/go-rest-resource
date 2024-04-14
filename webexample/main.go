package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/slyjeff/rest-resource/encoding"
	"github.com/slyjeff/rest-resource/resource"
	"net/http"
)

func main() {
	e := echo.New()
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(301, "/application")
	})

	e.GET("/application", func(c echo.Context) error {
		r := resource.NewResource("Application")
		r.Link("/self", "/application")
		r.LinkWithParameters("searchUsers", "/user").
			Parameter("username")

		return respond(c, r)
	})

	registerUserHandlers(e)

	e.Logger.Fatal(e.Start(":8090"))
}

func respond(c echo.Context, r resource.Resource) error {
	value, contentType := encoding.MarshalResource(c.Request().Header, r)
	c.Response().Header().Set("Content-Type", contentType)
	return c.String(http.StatusOK, value)
}
