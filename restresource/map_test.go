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
	intValue, ok := resource.Values["intValue"]
	a.True(ok, "'intValue' must exist")
	a.Equal(982, intValue, "'intValue' value must be '982'")

	var stringValue interface{}
	stringValue, ok = resource.Values["stringValue"]
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
	intValue, ok := resource.Values["intValue"]
	a.True(ok, "'intValue' must exist")
	a.Equal(982, intValue, "'intValue' value must be '982'")

	var stringValue interface{}
	stringValue, ok = resource.Values["stringValue"]
	a.True(ok, "'stringValue' must exist")
	a.Equal("Some test text.", stringValue, "'stringValue' value must be 'Some test text'.")

	var boolValue interface{}
	boolValue, ok = resource.Values["boolValue"]
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
	value, ok := resource.Values["floatValue"]
	a.True(ok, "'floatValue' must exist")

	var fd FormattedData
	fd, ok = value.(FormattedData)
	a.True(ok, "'floatValue' must be of type formatted data")

	a.Equal(982.4332, fd.Value, "'floatValue' value must be '4234.3982'.")
	a.Equal("982.43", fd.FormattedString(), "'floatValue' value  frmatted as string correctly.")
}

func Test_MapSliceFromMustMapFromIndicatedProperties(t *testing.T) {
	//arrange
	values := []struct {
		IntValue    int
		StringValue string
		BoolValue   bool
	}{{
		IntValue:    982,
		StringValue: "Some test text",
		BoolValue:   false,
	}, {
		IntValue:    123,
		StringValue: "Some other text",
		BoolValue:   false,
	}}

	testSlice := make([]interface{}, len(values))
	for i, v := range values {
		testSlice[i] = v
	}

	var resource Resource

	//act
	resource.MapSliceFrom("testSlice", testSlice).
		Map("IntValue").
		Map("StringValue")

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]interface{})
	a.True(ok, "'testSlice' must exist")

	var item1 ResourceMap
	item1, ok = slice[0].(ResourceMap)
	a.True(ok, "'item1' must exist")

	var intValue1 interface{}
	intValue1, ok = item1.Values["intValue"]
	a.True(ok, "'intValue1' must exist")
	a.Equal(982, intValue1, "'intValue1' value must be '982'")

	var stringValue1 interface{}
	stringValue1, ok = item1.Values["stringValue"]
	a.True(ok, "'stringValue1' must exist")
	a.Equal("Some test text", stringValue1, "'stringValue1' value must be 'Some test text'")

	_, ok = item1.Values["boolValue"]
	a.False(ok, "'boolValue' must not exist")

	var item2 ResourceMap
	item2, ok = slice[1].(ResourceMap)
	a.True(ok, "'item1' must exist")

	var intValue2 interface{}
	intValue2, ok = item2.Values["intValue"]
	a.True(ok, "'intValue2' must exist")
	a.Equal(123, intValue2, "'intValue2' value must be '123'")

	var stringValue2 interface{}
	stringValue2, ok = item2.Values["stringValue"]
	a.True(ok, "'stringValue2' must exist")
	a.Equal("Some other text", stringValue2, "'stringValue2' value must be 'Some other text'")
}

func Test_MapDataFromMustSupportMapAll(t *testing.T) {
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
	resource.MapDataFrom(testStruct).MapAll()

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

func Test_MapAllMustNotOverwriteFormattedString(t *testing.T) {
	//arrange
	testStruct := struct {
		FloatValue  float64
		StringValue string
		BoolValue   bool
	}{
		FloatValue:  982.43564,
		StringValue: "Some test text.",
		BoolValue:   false,
	}

	var resource Resource

	formatToTwoDecimals := func(v interface{}) string { return fmt.Sprintf("%.02f", v) }

	//act
	resource.MapDataFrom(testStruct).
		MapFormatted("FloatValue", formatToTwoDecimals).
		MapAll()

	//assert
	a := assert.New(t)
	value, _ := resource.Values["floatValue"].AsValue()
	var fd FormattedData
	fd, _ = value.(FormattedData)
	a.Equal("982.44", fd.FormattedString(), "'floatValue' value  formatted as string correctly.")
}

func Test_MapAllMustNotIncludeExcludedFields(t *testing.T) {
	//arrange
	testStruct := struct {
		FloatValue  float64
		StringValue string
		BoolValue   bool
	}{
		FloatValue:  982.43564,
		StringValue: "Some test text.",
		BoolValue:   false,
	}

	var resource Resource

	//act
	resource.MapDataFrom(testStruct).
		Exclude("FloatValue").
		MapAll().
		Exclude("BoolValue")

	//assert
	a := assert.New(t)
	var ok bool

	_, ok = resource.Values["floatValue"]
	a.False(ok, "floatValue must be excluded.")

	_, ok = resource.Values["stringValue"]
	a.True(ok, "stringValue must not be excluded.")

	_, ok = resource.Values["boolValue"]
	a.False(ok, "boolValue must be excluded.")
}

func Test_MapAllDataFromMustNotRequireEndMap(t *testing.T) {
	//arrange
	testStruct := struct {
		IntValue    int
		StringValue string
	}{
		IntValue:    982,
		StringValue: "Some test text.",
	}

	var resource Resource

	//act
	resource.MapAllDataFrom(testStruct).
		Data("boolValue", false)

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
