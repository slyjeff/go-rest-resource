package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/slyjeff/rest-resource/encoding"
	"github.com/slyjeff/rest-resource/resource"
	"net/http"
)

func main() {
	userRepo := newUserRepo()

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

	e.GET("/user", func(c echo.Context) error {
		userSearch := userSearch{}
		if err := c.Bind(&userSearch); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}

		users := userRepo.Search(userSearch.Username)

		r := newUserListResource(users, userSearch.Criteria())

		return respond(c, r)
	})

	e.POST("/user", func(c echo.Context) error {
		user := user{}
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}
		r := newUserResource(user)

		return respond(c, r)
	})

	e.Logger.Fatal(e.Start(":8090"))
}

func respond(c echo.Context, r resource.Resource) error {
	value, contentType := encoding.MarshalResource(c.Request().Header, r)
	c.Response().Header().Set("Content-Type", contentType)
	return c.String(http.StatusOK, value)
}
