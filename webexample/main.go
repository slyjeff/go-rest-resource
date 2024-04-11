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
		alecia := newUser("ajones", true, nil)
		susan := newUser("sanderson", false, nil)
		mark := newUser("mwilliams", false, nil)
		joe := newUser("jsmith", true, &alecia, susan, mark)

		userResource := newUserResource(joe)

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

func newUserResource(user user) resource.Resource {
	userResource := resource.NewResource()
	userResource.MapAllDataFrom(user)
	userResource.Data("lastPaycheck", 30324.2534, option.Format("%.02f"))
	return userResource
}

func respond(c echo.Context, r resource.Resource) error {
	value, contentType := encoding.MarshalResource(c.Request().Header, r)
	c.Response().Header().Set("Content-Type", contentType)
	return c.String(http.StatusOK, value)
}
