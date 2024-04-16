package main

import (
	"github.com/labstack/echo/v4"
	resource "github.com/slyjeff/rest-resource"
	openapi2 "github.com/slyjeff/rest-resource/openapi"
	"net/http"
)

func getDocumentation(c echo.Context) error {
	info := openapi2.Info{
		Title:   "Test Service",
		Version: "0.0.1",
	}

	defaultUser := user{}
	defaultUserList := []user{defaultUser}

	resources := []resource.Resource{
		newApplicationResource(),
		newUserListResource(defaultUserList, ""),
		newUserResource(defaultUser),
	}

	marshalledJson, contentType := openapi2.MarshalDoc(c.Request().Header, info, "http://localhost:8090/", resources...)

	c.Response().Header().Set("Content-Type", contentType)
	return c.String(http.StatusOK, string(marshalledJson))
}
