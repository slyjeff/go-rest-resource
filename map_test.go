package GoRestResource

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

func Test_MapDataFromSupportFormattedData(t *testing.T) {
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
		MapWithOptions("FloatValue", MapOptions{FormatCallback: formatToTwoDecimals})

	//assert
	a := assert.New(t)
	value, ok := resource.Values["floatValue"]
	a.True(ok, "'floatValue' must exist")

	var fd FormattedData
	fd, ok = value.(FormattedData)
	a.True(ok, "'floatValue' must be of type formatted data")

	a.Equal(982.4332, fd.Value, "'floatValue' value must be '4234.3982'.")
	a.Equal("982.43", fd.FormattedString(), "'floatValue' value  formatted as string correctly.")
}

func Test_MapDataFromSupportRenaming(t *testing.T) {
	//arrange
	testStruct := struct {
		Value string
	}{
		Value: "test value",
	}

	var resource Resource

	//act
	resource.MapDataFrom(testStruct).
		MapWithOptions("Value", MapOptions{Name: "coolValue"})

	//assert
	a := assert.New(t)
	value, ok := resource.Values["coolValue"]
	a.True(ok, "'coolValue' must exist")

	a.Equal("test value", value, "'coolValue' value must be 'test value'.")
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

func Test_MapAllMustNotOverwriteMapOptions(t *testing.T) {
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
		MapWithOptions("FloatValue", MapOptions{FormatCallback: formatToTwoDecimals}).
		MapAll()

	//assert
	a := assert.New(t)
	value, _ := resource.Values["floatValue"]
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

func Test_MapFromSliceMustMapFromIndicatedProperties(t *testing.T) {
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
	resource.MapFromSlice("testSlice", testSlice).
		Map("IntValue").
		Map("StringValue")

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var intValue1 interface{}
	intValue1, ok = slice[0].Values["intValue"]
	a.True(ok, "'intValue1' must exist")
	a.Equal(982, intValue1, "'intValue1' value must be '982'")

	var stringValue1 interface{}
	stringValue1, ok = slice[0].Values["stringValue"]
	a.True(ok, "'stringValue1' must exist")
	a.Equal("Some test text", stringValue1, "'stringValue1' value must be 'Some test text'")

	_, ok = slice[0].Values["boolValue"]
	a.False(ok, "'boolValue' must not exist")

	var intValue2 interface{}
	intValue2, ok = slice[1].Values["intValue"]
	a.True(ok, "'intValue2' must exist")
	a.Equal(123, intValue2, "'intValue2' value must be '123'")

	var stringValue2 interface{}
	stringValue2, ok = slice[1].Values["stringValue"]
	a.True(ok, "'stringValue2' must exist")
	a.Equal("Some other text", stringValue2, "'stringValue2' value must be 'Some other text'")
}

func Test_MapFromSliceMustAllowRenaming(t *testing.T) {
	//arrange
	values := []struct {
		IntValue int
	}{{
		IntValue: 13,
	}}

	testSlice := make([]interface{}, len(values))
	for i, v := range values {
		testSlice[i] = v
	}

	var resource Resource

	//act
	resource.MapFromSlice("testSlice", testSlice).
		MapWithOptions("IntValue", MapOptions{Name: "age"})

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var intValue1 interface{}
	intValue1, ok = slice[0].Values["age"]
	a.True(ok, "'age' must exist")
	a.Equal(13, intValue1, "'age' value must be '13'")
}

func Test_MapFromSliceMustAllowFormatting(t *testing.T) {
	//arrange
	values := []struct {
		FloatValue float64
	}{{
		FloatValue: 53.255,
	}}

	testSlice := make([]interface{}, len(values))
	for i, v := range values {
		testSlice[i] = v
	}

	var resource Resource

	formatToTwoDecimals := func(v interface{}) string { return fmt.Sprintf("%.02f", v) }

	//act
	resource.MapFromSlice("testSlice", testSlice).
		MapWithOptions("FloatValue", MapOptions{FormatCallback: formatToTwoDecimals})

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var floatValue interface{}
	floatValue, ok = slice[0].Values["floatValue"]
	a.True(ok, "'floatValue' must exist")

	var fd FormattedData
	fd, ok = floatValue.(FormattedData)
	a.True(ok, "'floatValue' must be of type formatted data")

	a.Equal(53.255, fd.Value, "'floatValue' value must be '4234.3982'.")
	a.Equal("53.26", fd.FormattedString(), "'floatValue' value formatted as string correctly.")
}

func Test_MapFromSliceMustSupportMapAll(t *testing.T) {
	//arrange
	values := []struct {
		IntValue    int
		StringValue string
	}{{
		IntValue:    982,
		StringValue: "Some test text",
	}, {
		IntValue:    123,
		StringValue: "Some other text",
	}}

	testSlice := make([]interface{}, len(values))
	for i, v := range values {
		testSlice[i] = v
	}

	var resource Resource

	//act
	resource.MapFromSlice("testSlice", testSlice).
		MapAll()

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var intValue1 interface{}
	intValue1, ok = slice[0].Values["intValue"]
	a.True(ok, "'intValue1' must exist")
	a.Equal(982, intValue1, "'intValue1' value must be '982'")

	var stringValue1 interface{}
	stringValue1, ok = slice[0].Values["stringValue"]
	a.True(ok, "'stringValue1' must exist")
	a.Equal("Some test text", stringValue1, "'stringValue1' value must be 'Some test text'")

	var intValue2 interface{}
	intValue2, ok = slice[1].Values["intValue"]
	a.True(ok, "'intValue2' must exist")
	a.Equal(123, intValue2, "'intValue2' value must be '123'")

	var stringValue2 interface{}
	stringValue2, ok = slice[1].Values["stringValue"]
	a.True(ok, "'stringValue2' must exist")
	a.Equal("Some other text", stringValue2, "'stringValue2' value must be 'Some other text'")
}

func Test_MapFromSliceMustNotOverwriteMapOptions(t *testing.T) {
	//arrange
	values := []struct {
		IntValue    int
		StringValue string
	}{{
		IntValue:    49,
		StringValue: "Some test text",
	}, {
		IntValue:    36,
		StringValue: "Some other text",
	}}

	testSlice := make([]interface{}, len(values))
	for i, v := range values {
		testSlice[i] = v
	}

	var resource Resource

	//act
	resource.MapFromSlice("testSlice", testSlice).
		MapWithOptions("IntValue", MapOptions{Name: "age"}).
		MapAll()

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var age1 interface{}
	age1, ok = slice[0].Values["age"]
	a.True(ok, "'age1' must exist")
	a.Equal(49, age1, "'age1' value must be '49'")

	var age2 interface{}
	age2, ok = slice[1].Values["intValue"]
	a.True(ok, "'age2' must exist")
	a.Equal(36, age2, "'intValue2' value must be '36'")
}

func Test_MapFromSliceMustMustNotIncludeExcludedFields(t *testing.T) {
	//arrange
	values := []struct {
		IntValue    int
		StringValue string
	}{{
		IntValue:    49,
		StringValue: "Some test text",
	}, {
		IntValue:    36,
		StringValue: "Some other text",
	}}

	testSlice := make([]interface{}, len(values))
	for i, v := range values {
		testSlice[i] = v
	}

	var resource Resource

	//act
	resource.MapFromSlice("testSlice", testSlice).
		Exclude("IntValue").
		MapAll()

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	_, ok = slice[0].Values["intValue"]
	a.False(ok, "'intValue' must not exist")

	var stringValue1 interface{}
	stringValue1, ok = slice[0].Values["stringValue"]
	a.True(ok, "'stringValue1' must exist")
	a.Equal("Some test text", stringValue1, "'stringValue1' value must be 'Some test text'")

	_, ok = slice[1].Values["intValue"]
	a.False(ok, "'intValue' must not exist")

	var stringValue2 interface{}
	stringValue2, ok = slice[1].Values["stringValue"]
	a.True(ok, "'age2' must exist")
	a.Equal("Some other text", stringValue2, "'stringValue2' value must be 'Some other text'")
}

func Test_MapAllFromSliceMustNotRequireEndMap(t *testing.T) {
	//arrange
	values := []struct {
		IntValue    int
		StringValue string
	}{{
		IntValue:    982,
		StringValue: "Some test text",
	}, {
		IntValue:    123,
		StringValue: "Some other text",
	}}

	testSlice := make([]interface{}, len(values))
	for i, v := range values {
		testSlice[i] = v
	}

	var resource Resource

	//act
	resource.MapAllDataFromSlice("testSlice", testSlice).
		Data("boolValue", false)

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["testSlice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var intValue1 interface{}
	intValue1, ok = slice[0].Values["intValue"]
	a.True(ok, "'intValue1' must exist")
	a.Equal(982, intValue1, "'intValue1' value must be '982'")

	var stringValue1 interface{}
	stringValue1, ok = slice[0].Values["stringValue"]
	a.True(ok, "'stringValue1' must exist")
	a.Equal("Some test text", stringValue1, "'stringValue1' value must be 'Some test text'")

	var intValue2 interface{}
	intValue2, ok = slice[1].Values["intValue"]
	a.True(ok, "'intValue2' must exist")
	a.Equal(123, intValue2, "'intValue2' value must be '123'")

	var stringValue2 interface{}
	stringValue2, ok = slice[1].Values["stringValue"]
	a.True(ok, "'stringValue2' must exist")
	a.Equal("Some other text", stringValue2, "'stringValue2' value must be 'Some other text'")

	var boolValue interface{}
	boolValue, ok = resource.Values["boolValue"]
	a.True(ok, "'boolValue' must exist")
	a.Equal(false, boolValue, "'boolValue' value must be false.")
}

func Test_MustBeAbleToMapSliceMappingAStruct(t *testing.T) {
	//arrange
	values := []struct {
		IntValue    int
		StringValue string
	}{{
		IntValue:    982,
		StringValue: "Some test text",
	}, {
		IntValue:    123,
		StringValue: "Some other text",
	}}

	testStruct := struct {
		StringValue string
		IntValue    int
		Slice       []interface{}
	}{
		StringValue: "Hi there",
		IntValue:    9382,
		Slice:       make([]interface{}, len(values)),
	}

	for i, v := range values {
		testStruct.Slice[i] = v
	}

	var resource Resource

	//act
	resource.MapDataFrom(testStruct).
		Map("StringValue").
		Map("IntValue").
		Map("Slice").
		EndMap()

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["slice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var intValue1 interface{}
	intValue1, ok = slice[0].Values["intValue"]
	a.True(ok, "'intValue1' must exist")
	a.Equal(982, intValue1, "'intValue1' value must be '982'")

	var stringValue1 interface{}
	stringValue1, ok = slice[0].Values["stringValue"]
	a.True(ok, "'stringValue1' must exist")
	a.Equal("Some test text", stringValue1, "'stringValue1' value must be 'Some test text'")

	var intValue2 interface{}
	intValue2, ok = slice[1].Values["intValue"]
	a.True(ok, "'intValue2' must exist")
	a.Equal(123, intValue2, "'intValue2' value must be '123'")

	var stringValue2 interface{}
	stringValue2, ok = slice[1].Values["stringValue"]
	a.True(ok, "'stringValue2' must exist")
	a.Equal("Some other text", stringValue2, "'stringValue2' value must be 'Some other text'")
}

func Test_MustBeAbleToMapSliceByFieldWhenMappingAStruct(t *testing.T) {
	//arrange
	values := []struct {
		IntValue    int
		StringValue string
	}{{
		IntValue:    982,
		StringValue: "Some test text",
	}, {
		IntValue:    123,
		StringValue: "Some other text",
	}}

	testStruct := struct {
		StringValue string
		IntValue    int
		Slice       []interface{}
	}{
		StringValue: "Hi there",
		IntValue:    9382,
		Slice:       make([]interface{}, len(values)),
	}

	for i, v := range values {
		testStruct.Slice[i] = v
	}

	var resource Resource

	//act
	resource.MapDataFrom(testStruct).
		Map("StringValue").
		Map("IntValue").
		MapChild("Slice").
		Map("StringValue").
		Map("IntValue").
		EndMap().
		EndMap()

	//assert
	a := assert.New(t)
	slice, ok := resource.Values["slice"].([]ResourceMap)
	a.True(ok, "'testSlice' must exist")

	var intValue1 interface{}
	intValue1, ok = slice[0].Values["intValue"]
	a.True(ok, "'intValue1' must exist")
	a.Equal(982, intValue1, "'intValue1' value must be '982'")

	var stringValue1 interface{}
	stringValue1, ok = slice[0].Values["stringValue"]
	a.True(ok, "'stringValue1' must exist")
	a.Equal("Some test text", stringValue1, "'stringValue1' value must be 'Some test text'")

	var intValue2 interface{}
	intValue2, ok = slice[1].Values["intValue"]
	a.True(ok, "'intValue2' must exist")
	a.Equal(123, intValue2, "'intValue2' value must be '123'")

	var stringValue2 interface{}
	stringValue2, ok = slice[1].Values["stringValue"]
	a.True(ok, "'stringValue2' must exist")
	a.Equal("Some other text", stringValue2, "'stringValue2' value must be 'Some other text'")
}
