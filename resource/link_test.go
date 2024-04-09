package resource

import (
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
	a.True(ok, "'getUser' must exist")
	a.Equal(link.Href, href)
	a.Equal(link.Verb, "GET")
}