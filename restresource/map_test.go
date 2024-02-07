package restresource

import (
	"fmt"
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

func Test_MapDataFromMustAddFromMultipleStructs(t *testing.T) {
	//arrange
	testStruct := struct {
		IntValue    int
		StringValue string
	}{
		IntValue:    982,
		StringValue: "Some test text.",
	}

	testStruct2 := struct {
		BoolValue bool
	}{
		BoolValue: false,
	}

	var resource Resource

	//act
	resource.MapDataFrom(testStruct).
		Map("IntValue").
		Map("StringValue").
		EndMap().
		MapDataFrom(testStruct2).
		Map("BoolValue")

	//assert
	a := assert.New(t)
	intValue, ok := resource.Values["intValue"].AsValue()
	a.True(ok, "'intValue' must exist")
	a.Equal(982, intValue, "'intValue' value must be '982'")

	var stringValue interface{}
	stringValue, ok = resource.Values["stringValue"].AsValue()
	a.True(ok, "'stringValue' must exist")
	a.Equal("Some test text.", stringValue, "'stringValue' value must be 'Some test text'.")

	var boolValue interface{}
	boolValue, ok = resource.Values["boolValue"].AsValue()
	a.True(ok, "'boolValue' must exist")
	a.Equal(false, boolValue, "'boolValue' value must be false.")
}

func Test_MapDataFromMustAddFormattedData(t *testing.T) {
	//arrange
	testStruct := struct {
		FloatValue float64
	}{
		FloatValue: 982.4332,
	}

	var resource Resource

	formatToTwoDecimals := func(v interface{}) string { return fmt.Sprintf("%.02f", v) }

	//act
	resource.MapDataFrom(testStruct).
		MapFormatted("FloatValue", formatToTwoDecimals)

	//assert
	a := assert.New(t)
	value, ok := resource.Values["floatValue"].AsValue()
	a.True(ok, "'floatValue' must exist")

	var fd FormattedData
	fd, ok = value.(FormattedData)
	a.True(ok, "'floatValue' must be of type formatted data")

	a.Equal(982.4332, fd.Value, "'floatValue' value must be '4234.3982'.")
	a.Equal("982.43", fd.FormattedString(), "'floatValue' value  formatted as string correctly.")
}
