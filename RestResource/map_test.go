package restresource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MapDataFromMustAddIndicatedProperties(t *testing.T) {
	//arrange
	testStruct := struct {
		IntValue    int
		StringValue string
		BoolValue   bool
	}{
		IntValue:    982,
		StringValue: "Some test text.",
		BoolValue:   false,
	}

	var resource Resource

	//act
	resource.MapDataFrom(testStruct).
		Map("IntValue").
		Map("StringValue")

	//assert
	a := assert.New(t)
	intValue, ok := resource.Values["intValue"].AsValue()
	a.True(ok, "'intValue' must exist")
	a.Equal(982, intValue, "'intValue' value must be '982'")

	var stringValue interface{}
	stringValue, ok = resource.Values["stringValue"].AsValue()
	a.True(ok, "'stringValue' must exist")
	a.Equal("Some test text.", stringValue, "'stringValue' value must be 'Some test text.'")

	_, ok = resource.Values["boolValue"]
	a.False(ok, "'boolValue' must not exist")
}
