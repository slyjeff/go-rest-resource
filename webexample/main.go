package main

import (
	"github.com/labstack/echo/v4"
	"github.com/slyjeff/rest-resource/encoding"
	"github.com/slyjeff/rest-resource/resource"
	"github.com/slyjeff/rest-resource/resource/mapping"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		alecia := newUser("ajones", true, nil)
		susan := newUser("sanderson", false, nil)
		mark := newUser("mwilliams", false, nil)
		joe := newUser("jsmith", true, &alecia, susan, mark)

		var userResource resource.Resource
		userResource.MapAllDataFrom(joe)
		userResource.Data("lastPaycheck", 30324.2534, mapping.Format("%.02f"))

		return respond(c, userResource)
	})
	e.Logger.Fatal(e.Start(":80"))
}

type user struct {
	Username      string
	IsAdmin       bool
	Supervisor    *user
	DirectReports []user
}

func newUser(userName string, isAdmin bool, supervisor *user, directReports ...user) user {
	return user{
		userName,
		isAdmin,
		supervisor,
		directReports,
	}
}

func respond(c echo.Context, r resource.Resource) error {
	value, contentType := encoding.MarshalResource(r, c.Request().Header)
	c.Response().Header().Set("Content-Type", contentType)
	return c.String(http.StatusOK, value)
}
