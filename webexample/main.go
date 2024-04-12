package main

import (
	"github.com/labstack/echo/v4"
	"github.com/slyjeff/rest-resource/encoding"
	"github.com/slyjeff/rest-resource/openapi"
	"github.com/slyjeff/rest-resource/resource"
	"github.com/slyjeff/rest-resource/resource/option"
	"net/http"
)

func main() {
	//userRepo := newUserRepo()

	e := echo.New()
	e.GET("/doc", func(c echo.Context) error {
		info := openapi.Info{
			Title:   "A Test API",
			Version: "1.0.0",
		}

		userResource := newUserResource(user{})
		value, contentType := openapi.MarshalDoc(c.Request().Header, info, userResource)
		c.Response().Header().Set("Content-Type", contentType)
		return c.String(http.StatusOK, value)
	})

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(301, "/application")
	})

	e.GET("/application", func(c echo.Context) error {
		r := resource.NewResource("Application")
		r.Link("getUsers", "/user")
		return respond(c, r)
	})

	e.GET("/user", func(c echo.Context) error {
		r := resource.NewResource("Users")
		r.Data("message", "Here be users")
		r.LinkWithParameters("addUser", "/user", option.Verb("POST")).
			Parameter("userName").
			Parameter("email")

		return respond(c, r)
	})

	e.Logger.Fatal(e.Start(":8090"))
}

func newUserResource(user user) resource.Resource {
	userResource := resource.NewResource("User")
	userResource.MapAllDataFrom(user)
	return userResource
}

func respond(c echo.Context, r resource.Resource) error {
	value, contentType := encoding.MarshalResource(c.Request().Header, r)
	c.Response().Header().Set("Content-Type", contentType)
	return c.String(http.StatusOK, value)
}
