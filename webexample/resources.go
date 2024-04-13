package main

import (
	"github.com/slyjeff/rest-resource/resource"
	"github.com/slyjeff/rest-resource/resource/option"
)

func newUserListResource(users []user, queryParams string) resource.Resource {
	userListResource := resource.NewResource("UserList")

	userResources := make([]resource.Resource, len(users))
	for i, user := range users {
		userResources[i] = newUserResource(user)
	}

	userListResource.
		Embed("users", userResources...).
		Link("_self", "/user"+queryParams).
		LinkWithParameters("createUser", "/user", option.Verb("POST")).
		Parameter("userName").
		Parameter("Email")

	return userListResource
}

func newUserResource(user user) resource.Resource {
	userResource := resource.NewResource("User")
	userResource.MapAllDataFrom(user)
	return userResource
}
