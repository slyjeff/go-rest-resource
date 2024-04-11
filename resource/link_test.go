package resource

import (
	"github.com/slyjeff/rest-resource/resource/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LinkMustAddLinkToResource(t *testing.T) {
	//arrange
	href := "/user"
	var resource Resource

	//act
	resource.Link("getUser", href)

	//assert
	a := assert.New(t)
	link, ok := resource.Links["getUser"]
	a.True(ok)
	a.Equal(link.Href, href)
	a.Equal(link.Verb, "GET")
	a.False(link.IsTemplated)
}

func Test_LinkMustAddLinkToResourceWithVerb(t *testing.T) {
	//arrange
	var resource Resource

	//act
	resource.Link("newUser", "/user", option.Verb("POST"))

	//assert
	a := assert.New(t)
	link, _ := resource.Links["newUser"]
	a.Equal(link.Verb, "POST")
}

func Test_LinkMustAddLinkWithTemplated(t *testing.T) {
	//arrange
	var resource Resource

	//act
	resource.Link("getUser", "/user/{id}", option.Templated())

	//assert
	a := assert.New(t)
	link, _ := resource.Links["getUser"]
	a.True(link.IsTemplated)
}

func Test_LinkMustAddLinkToResourceWithParameters(t *testing.T) {
	//arrange
	var resource Resource

	//act
	resource.LinkWithParameters("searchUsers", "/user").
		Parameter("lastName").
		Parameter("firstName").
		EndMap().
		Link("newUser", "/user", option.Verb("POST"))

	//assert
	a := assert.New(t)
	link, _ := resource.Links["searchUsers"]
	a.Equal(link.Href, "/user")
	a.Equal(link.Verb, "GET")
	a.Equal(link.Parameters[0].Name, "lastName")
	a.Equal(link.Parameters[1].Name, "firstName")

	_, ok := resource.Links["newUser"]
	a.True(ok)
}

func Test_LinkMustAddParametersWithDefaultValues(t *testing.T) {
	//arrange
	var resource Resource

	//act
	resource.LinkWithParameters("searchUsers", "/user").
		Parameter("lastName", option.Default("Smith")).
		Parameter("firstName")

	//assert
	a := assert.New(t)
	link, _ := resource.Links["searchUsers"]
	a.Equal(link.Href, "/user")
	a.Equal(link.Verb, "GET")
	a.Equal(link.Parameters[0].DefaultValue, "Smith")
}
