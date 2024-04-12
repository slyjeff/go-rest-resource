package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewResourceWithoutSettingNameMustBeEmpty(t *testing.T) {
	//arrange
	//act
	resource := NewResource()

	//assert
	a := assert.New(t)
	a.Equal(resource.Name, "")
}

func Test_NewResourceMustSetName(t *testing.T) {
	//arrange
	name := "Test Name"

	//act
	resource := NewResource(name)

	//assert
	a := assert.New(t)
	a.Equal(resource.Name, name)
}
