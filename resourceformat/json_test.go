package resourceformat

import (
	"github.com/stretchr/testify/assert"
	"restresource/restresource"
	"testing"
)

func Test_DataMustAddStringToResource(t *testing.T) {
	//arrange
	stringValue := "I am a string"
	intValue := 938

	var resource restresource.Resource
	resource.
		Data("stringValue", stringValue).
		Data("intValue", intValue)

	//act
	json := ToSlySoftHalJson(resource)

	//assert
	a := assert.New(t)
	expectedJson := "{\n" +
		"  \"stringValue\":\"I am a string\",\n" +
		"  \"intValue\":938\n" +
		"}"
	a.Equal(expectedJson, json, "json not created properly.")
}
