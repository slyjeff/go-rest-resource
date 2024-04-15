package encoding

import (
	"fmt"
	"github.com/slyjeff/rest-resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func validateResourceData(a *assert.Assertions, dataName string, originalResource, unmarshalledResource resource.Resource) {
	originalValue, ok := originalResource.Values[dataName]
	if !ok {
		a.Fail(fmt.Sprintf("expected '%s' to exist in original resource", dataName))
		return
	}

	var newValue interface{}
	newValue, ok = unmarshalledResource.Values[dataName]
	if !ok {
		a.Fail(fmt.Sprintf("expected '%s' to exist in unmarshalled resource", dataName))
		return
	}

	a.Equal(fmt.Sprintf("%v", originalValue), fmt.Sprintf("%v", newValue))
}

func Test_UnmarshalJsonMustDecodeAllProperties(t *testing.T) {
	//arrange
	testItem := newTestItem1()

	originalResource := resource.NewResource()
	originalResource.MapAllDataFrom(testItem)
	json, _ := MarshalJson(originalResource)

	//act
	unmarshalledResource, err := UnmarshalJson(json)

	//assert
	a := assert.New(t)
	a.NoError(err)
	validateResourceData(a, "IsAvailable", originalResource, unmarshalledResource)
	validateResourceData(a, "Name", originalResource, unmarshalledResource)
	validateResourceData(a, "Price", originalResource, unmarshalledResource)
	validateResourceData(a, "Quantity", originalResource, unmarshalledResource)
}

func Test_UnmarshalJsonMustReturnErrorForInvalidJson(t *testing.T) {
	//arrange
	invalidJson := make([]byte, 0)

	//act
	_, err := UnmarshalJson(invalidJson)

	//assert
	a := assert.New(t)
	a.Error(err)
}

func Test_UnmarshalJsonMustDecodeLink(t *testing.T) {
	//arrange
	originalResource := resource.NewResource()
	originalResource.Link("self", "/user")
	json, _ := MarshalJson(originalResource)

	//act
	unmarshalledResource, err := UnmarshalJson(json)

	//assert
	a := assert.New(t)
	a.NoError(err)

	_, ok := unmarshalledResource.Values["_links"]
	a.False(ok, "_links shouldn't be added to values")

	var link *resource.Link
	link, ok = unmarshalledResource.Links["self"]
	a.True(ok)
	a.Equal("/user", link.Href)
}
