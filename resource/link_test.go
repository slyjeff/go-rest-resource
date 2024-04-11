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
	a.Equal("GET", link.Verb)
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
	a.Equal("POST", link.Verb)
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
	a.Equal("lastName", link.Parameters[0].Name)
	a.Equal("firstName", link.Parameters[1].Name)

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
	a.Equal("Smith", link.Parameters[0].DefaultValue)
}

func Test_LinkMustAddParametersWithListOfValues(t *testing.T) {
	//arrange
	var resource Resource
	var values = []int{1, 2, 3}

	//act
	resource.LinkWithParameters("searchUsers", "/user").
		Parameter("lastName", option.ListOfValues(values)).
		Parameter("firstName")

	//assert
	a := assert.New(t)
	link, _ := resource.Links["searchUsers"]
	a.Equal(link.Href, "/user")
	a.Equal(link.Verb, "GET")
	a.Equal("1,2,3", link.Parameters[0].ListOfValues)
}
