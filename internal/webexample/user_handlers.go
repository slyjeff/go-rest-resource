package main

import (
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

		users := userRepo.Search(userSearch.Username)

		r := newUserListResource(users, userSearch.Criteria())

		return respond(c, r)
	})

	e.GET("/user/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		u, ok := userRepo.GetById(id)
		if !ok {
			return c.String(http.StatusNotFound, "User not found.")
		}

		r := newUserResource(*u)
		return respond(c, r)
	})

	e.POST("/user", func(c echo.Context) error {
		user := user{}
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}
		userRepo.Add(&user)

		r := newUserResource(user)

		return respond(c, r)
	})

	e.PUT("/user/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		u, ok := userRepo.GetById(id)
		if !ok {
			return c.String(http.StatusNotFound, "User not found.")
		}

		if err := c.Bind(u); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}

		r := newUserResource(*u)

		return respond(c, r)
	})

	e.DELETE("/user/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
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
}

func (s userSearch) Criteria() string {
	if s.Username == "" {
		return ""
	}
	return "?username=" + s.Username
}

func newUserListResource(users []user, queryParams string) resource.Resource {
	r := resource.NewResource("UserList")
	r.Uri("/user" + queryParams)

	userResources := make([]resource.Resource, len(users))
	for i, user := range users {
		userResources[i] = newUserResource(user)
	}

	r.EmbedResources("users", userResources)
	r.LinkWithParameters("createUser", "/user", option.Verb("POST")).
		Parameter("userName").
		Parameter("Email")

	return r
}

func newUserResource(user user) resource.Resource {
	url := resource.ConstructUriFromTemplate("/user/{id}", user.id)
	r := resource.NewResource("User")
	r.Uri(url)
	r.MapAllDataFrom(user)
	r.LinkWithParameters("updateUser", url, option.Verb("PUT")).
		Parameter("username", option.Default(user.Username)).
		Parameter("email", option.Default(user.Email))
	r.Link("deleteUser", url, option.Verb("DELETE"))
	return r
}
