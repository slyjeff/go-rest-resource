package resource

import (
	"github.com/slyjeff/rest-resource/option"
	"github.com/stretchr/testify/assert"
	"net/http"
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

func Test_UriMustAddSelfLinkToResource(t *testing.T) {
	//arrange
	resource := NewResource()

	//act
	resource.Uri("/user")

	//assert
	a := assert.New(t)
	link, ok := resource.Links["self"]
	a.True(ok)
	a.Equal("/user", link.Href)
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
	r := NewResource()

	//act
	r.Link("searchUsers", "/user").
		Parameter("lastName").
		Parameter("firstName")
	r.Link("newUser", "/user", option.Verb("POST"))

	//assert
	a := assert.New(t)
	link, _ := r.Links["searchUsers"]
	a.Equal(link.Href, "/user")
	a.Equal(link.Verb, "GET")
	a.Equal("lastName", link.Parameters[0].Name)
	a.Equal("firstName", link.Parameters[1].Name)

	_, ok := r.Links["newUser"]
	a.True(ok)
}

func Test_LinkMustAddParametersWithDefaultValues(t *testing.T) {
	//arrange
	var resource Resource

	//act
	resource.Link("searchUsers", "/user").
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
	resource.Link("searchUsers", "/user").
		Parameter("lastName", option.ListOfValues(values)).
		Parameter("firstName")

	//assert
	a := assert.New(t)
	link, _ := resource.Links["searchUsers"]
	a.Equal(link.Href, "/user")
	a.Equal(link.Verb, "GET")
	a.Equal("1,2,3", link.Parameters[0].ListOfValues)
}

func Test_LinkMustSetSchema(t *testing.T) {
	//arrange
	r := NewResource("User")

	//act
	r.Link("searchUsers", "/user").
		Schema("UserList")

	//assert
	a := assert.New(t)
	link, _ := r.Links["searchUsers"]

	a.Equal("UserList", link.Schema)
}

func Test_LinkMustSetResponseCodes(t *testing.T) {
	//arrange
	r := NewResource("User")

	//act
	r.Link("searchUsers", "/user").
		ResponseCodes(http.StatusOK, http.StatusInternalServerError)

	//assert
	a := assert.New(t)
	link, _ := r.Links["searchUsers"]

	a.Equal(http.StatusOK, link.ResponseCodes[0])
	a.Equal(http.StatusInternalServerError, link.ResponseCodes[1])
}
