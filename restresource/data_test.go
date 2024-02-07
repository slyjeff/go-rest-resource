package restresource

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DataMustAddStringToResource(t *testing.T) {
	//arrange
	message := "Test Message"
	var resource Resource

	//act
	resource.Data("message", message)

	//assert
	a := assert.New(t)
	value, ok := resource.Values["message"].AsValue()
	a.True(ok, "'message' must exist")
	a.Equal(message, value, "'message' value must be 'TestMessage'")

	_, ok = resource.Values["message"].AsSlice()
	a.False(ok, "'message' must not be a slice")

	_, ok = resource.Values["message"].AsMap()
	a.False(ok, "'message' must not be a map")
}

func Test_DataNameMustBeCamelCase(t *testing.T) {
	//arrange
	message := "Test Message"
	var resource Resource

	//act
	resource.Data("Message", message)

	//assert
	a := assert.New(t)
	_, ok := resource.Values["message"]
	a.True(ok, resource.Values["message"], "'message' name must start with a lowercase letter.")
}

func Test_DataMustStoreInt(t *testing.T) {
	//arrange
	number := 42
	var resource Resource

	//act
	resource.Data("number", number)

	//assert
	a := assert.New(t)
	value, ok := resource.Values["number"].AsValue()
	a.True(ok, "'number' must exist")
	a.Equal(42, value, "'number' value must be '42'.")
}

func Test_FormattedDataAddValueAndFormattingInformation(t *testing.T) {
	//arrange
	number := 4234.3982
	var resource Resource
	formatToTwoDecimals := func(v interface{}) string { return fmt.Sprintf("%.02f", v) }

	//act
	resource.FormattedData("number", number, formatToTwoDecimals)

	//assert
	a := assert.New(t)
	value, ok := resource.Values["number"].AsValue()
	a.True(ok, "'number' must exist")

	var fd FormattedData
	fd, ok = value.(FormattedData)
	a.True(ok, "'number' must be of type formatted data")

	a.Equal(4234.3982, fd.Value, "'number' value must be '4234.3982'.")
	a.Equal("4234.40", fd.FormattedString(), "'number' value  formatted as string correctly.")
}

func Test_DataMustBeChainable(t *testing.T) {
	//arrange
	value1 := 37
	value2 := "Some Text"
	var resource Resource

	//act
	resource.Data("value1", value1).
		Data("value2", value2)

	//assert
	a := assert.New(t)

	v1, ok := resource.Values["value1"].AsValue()
	a.True(ok, "'value1' must exist")
	a.Equal(37, v1, "'value1' value must be '37'.")

	var v2 interface{}
	v2, ok = resource.Values["value2"].AsValue()
	a.True(ok, "'value2' must exist")
	a.Equal("Some Text", v2, "'value2' value must be 'Some text'.")
}

func Test_DataMustTransformStructToMap(t *testing.T) {
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
	resource.Data("testStruct", testStruct)

	//assert
	a := assert.New(t)
	testStructAsMap, ok := resource.Values["testStruct"].AsMap()
	a.True(ok, "'testStruct' must be found in values.")

	var intValue interface{}
	intValue, ok = testStructAsMap["intValue"].AsValue()

	a.True(ok, "'intValue' must be int 'testStruct'.")
	a.Equal(982, intValue, "'intValue' value must be '982'.")

	var stringValue interface{}
	stringValue, ok = testStructAsMap["stringValue"].AsValue()

	a.True(ok, "'stringValue' must be int 'testStruct'.")
	a.Equal("Some test text.", stringValue, "'stringValue' value must be 'Some text'.")

	_, ok = resource.Values["testStruct"].AsValue()
	a.False(ok, "'testStruct' must not be a string")

	_, ok = resource.Values["testStruct"].AsSlice()
	a.False(ok, "'testStruct' must not be a slice")
}

func Test_DataMustAddSliceToResource(t *testing.T) {
	//arrange
	strings := []string{"text 1", "text 2", "text 3"}
	var resource Resource

	//act
	resource.Data("strings", strings)

	//assert
	a := assert.New(t)
	testStrings, ok := resource.Values["strings"].AsSlice()
	a.True(ok, "'strings' must be found in values.")

	var string1 interface{}
	string1, ok = testStrings[0].AsValue()
	a.True(ok, "element 0 must be a string.")
	a.Equal("text 1", string1, "element 0 must be 'text 1'.")

	var string2 interface{}
	string2, ok = testStrings[1].AsValue()
	a.True(ok, "element 1 must be a string.")
	a.Equal("text 2", string2, "element 1 must be 'text 2'.")

	var string3 interface{}
	string3, ok = testStrings[2].AsValue()
	a.True(ok, "element 2 must be a string.")
	a.Equal("text 3", string3, "element 2 must be 'text 3'.")

	_, ok = resource.Values["strings"].AsValue()
	a.False(ok, "'strings' must not be a string")

	_, ok = resource.Values["strings"].AsMap()
	a.False(ok, "'strings' must not be a map")
}

func Test_DataMustAddArrayToResource(t *testing.T) {
	//arrange
	strings := [...]string{"text 1", "text 2", "text 3"}
	var resource Resource

	//act
	resource.Data("strings", strings)

	//assert
	a := assert.New(t)
	testStrings, ok := resource.Values["strings"].AsSlice()
	a.True(ok, "'strings' must be found in values.")

	var string1 interface{}
	string1, ok = testStrings[0].AsValue()
	a.True(ok, "element 0 must be a string.")
	a.Equal("text 1", string1, "element 0 must be 'text 1'.")

	var string2 interface{}
	string2, ok = testStrings[1].AsValue()
	a.True(ok, "element 1 must be a string.")
	a.Equal("text 2", string2, "element 1 must be 'text 2'.")

	var string3 interface{}
	string3, ok = testStrings[2].AsValue()
	a.True(ok, "element 2 must be a string.")
	a.Equal("text 3", string3, "element 2 must be 'text 3'.")
}

func Test_DatMustTransformStructsToArraysInSlices(t *testing.T) {
	//arrange
	type testStruct struct {
		IntValue    int
		StringValue string
	}

	testStructs := []testStruct{
		{IntValue: 43, StringValue: "test 1"},
		{IntValue: 367, StringValue: "test 2"},
	}

	var resource Resource

	//act
	resource.Data("structs", testStructs)

	//assert
	a := assert.New(t)
	testMaps, ok := resource.Values["structs"].AsSlice()
	a.True(ok, "'structs' must be found in values.")

	var map1 map[string]ResourceData
	map1, ok = testMaps[0].AsMap()
	a.True(ok, "element 0 must be a map.")

	var intValue1 interface{}
	intValue1, ok = map1["intValue"].AsValue()
	a.True(ok, "'intValue' must be found in first testMap.")
	a.Equal(43, intValue1, "'intValue' must be '43'.")

	var stringValue1 interface{}
	stringValue1, ok = map1["stringValue"].AsValue()
	a.True(ok, "'stringValue' must be found in first testMap.")
	a.Equal("test 1", stringValue1, "'stringValue' must be 'test 1'.")

	var map2 map[string]ResourceData
	map2, ok = testMaps[1].AsMap()
	a.True(ok, "element 1 must be a map.")

	var intValue2 interface{}
	intValue2, ok = map2["intValue"].AsValue()
	a.True(ok, "'intValue' must be found in second testMap.")
	a.Equal(367, intValue2, "'intValue' must be '367'.")

	var stringValue2 interface{}
	stringValue2, ok = map2["stringValue"].AsValue()
	a.True(ok, "'stringValue' must be found in second testMap.")
	a.Equal("test 2", stringValue2, "'stringValue' must be 'test 2'.")
}
