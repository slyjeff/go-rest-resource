package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/slyjeff/rest-resource"
	"github.com/slyjeff/rest-resource/option"
	"net/http"
	"strconv"
)

func registerUserHandlers(e *echo.Echo) {
	userRepo := newUserRepo()

	e.GET("/user", func(c echo.Context) error {
		userSearch := userSearch{}
		if err := c.Bind(&userSearch); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}

		users := userRepo.Search(userSearch)

		r := newUserListResource(users, userSearch.Criteria())

		return respond(c, http.StatusOK, r)
	})

	e.GET("/user/:Id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("Id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Id")
		}

		u, ok := userRepo.GetById(id)
		if !ok {
			return c.String(http.StatusNotFound, "User not found.")
		}

		r := newUserResource(*u)
		return respond(c, http.StatusOK, r)
	})

	e.POST("/user", func(c echo.Context) error {
		user := user{}
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}
		userRepo.Add(&user)

		r := newUserResource(user)

		return respond(c, http.StatusCreated, r)
	})

	e.PUT("/user/:Id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("Id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Id")
		}

		u, ok := userRepo.GetById(id)
		if !ok {
			return c.String(http.StatusNotFound, "User not found.")
		}

		if err := c.Bind(u); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}

		r := newUserResource(*u)

		return respond(c, http.StatusOK, r)
	})

	e.DELETE("/user/:Id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("Id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Id")
		}

		ok := userRepo.Delete(id)
		if !ok {
			return c.String(http.StatusNotFound, "User not found.")
		}

		return c.String(http.StatusOK, "User deleted.")
	})
}

type userSearch struct {
	Username string `query:"username"`
	IsActive string `query:"is_active"`
}

func (s userSearch) Criteria() string {
	c := ""
	c = addQueryParameter(c, "username", s.Username)
	c = addQueryParameter(c, "is_active", s.IsActive)

	return c
}

func addQueryParameter(s, name string, value string) string {
	if name == "" || value == "" {
		return s
	}

	if s == "" {
		s += "?"
	} else {
		s += "&"
	}
	s += fmt.Sprintf("%s=%v", name, value)
	return s
}

func newUserListResource(users []user, queryParams string) resource.Resource {
	r := resource.NewResource("UserList")
	r.Uri("/user" + queryParams)

	userResources := make([]resource.Resource, len(users))
	for i, user := range users {
		userResources[i] = newUserResource(user)
	}

	r.EmbedResources("users", userResources)
	r.Link("createUser", "/user", option.Verb("POST")).
		Parameter("userName").
		Parameter("Email").
		ResponseSchema("User")

	return r
}

func newUserResource(user user) resource.Resource {
	url := resource.ConstructUriFromTemplate("/user/{Id}", user.Id)
	r := resource.NewResource("User")
	r.Uri(url)
	r.MapAllDataFrom(user)
	r.Link("updateUser", url, option.Verb("PUT")).
		Parameter("username", option.Default(user.Username)).
		Parameter("email", option.Default(user.Email)).
		ResponseSchema("User")
	r.Link("deleteUser", url, option.Verb("DELETE"))

	return r
}
