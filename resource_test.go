package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewResourceWithoutSettingSchemaMustBeEmpty(t *testing.T) {
	//arrange
	//act
	resource := NewResource()

	//assert
	a := assert.New(t)
	a.Equal(resource.Schema, "")
}

func Test_NewResourceMustSetSchema(t *testing.T) {
	//arrange
	schema := "TestSchema"

	//act
	resource := NewResource(schema)

	//assert
	a := assert.New(t)
	a.Equal(resource.Schema, schema)
}
