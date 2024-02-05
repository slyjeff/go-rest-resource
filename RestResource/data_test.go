package RestResource

import (
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
	value, ok := resource.Values["message"].AsString()
	a.True(ok, "'message' must exist")
	a.Equal(message, value, "'message' value must be 'TestMessage'.")

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

func Test_DataMustStoreIntAsStrings(t *testing.T) {
	//arrange
	number := 42
	var resource Resource

	//act
	resource.Data("number", number)

	//assert
	a := assert.New(t)
	value, ok := resource.Values["number"].AsString()
	a.True(ok, "'number' must exist")
	a.Equal("42", value, "'number' value must be '42'.")
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

	value1AsString, ok := resource.Values["value1"].AsString()
	a.True(ok, "'value1' must exist")
	a.Equal("37", value1AsString, "'value1' value must be '37'.")

	var value2AsString string
	value2AsString, ok = resource.Values["value2"].AsString()
	a.True(ok, "'value2' must exist")
	a.Equal("Some Text", value2AsString, "'value2' value must be 'Some text'.")
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

	var intValueAsString string
	intValueAsString, ok = testStructAsMap["intValue"].AsString()

	a.True(ok, "'intValue' must be int 'testStruct'.")
	a.Equal("982", intValueAsString, "'intValue' value must be '982'.")

	var stringValueAsString string
	stringValueAsString, ok = testStructAsMap["stringValue"].AsString()

	a.True(ok, "'stringValue' must be int 'testStruct'.")
	a.Equal("Some test text.", stringValueAsString, "'stringValue' value must be 'Some text'.")

	_, ok = resource.Values["testStruct"].AsString()
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

	var string1 string
	string1, ok = testStrings[0].AsString()
	a.True(ok, "element 0 must be a string.")
	a.Equal("text 1", string1, "element 0 must be 'text 1'.")

	var string2 string
	string2, ok = testStrings[1].AsString()
	a.True(ok, "element 1 must be a string.")
	a.Equal("text 2", string2, "element 1 must be 'text 2'.")

	var string3 string
	string3, ok = testStrings[2].AsString()
	a.True(ok, "element 2 must be a string.")
	a.Equal("text 3", string3, "element 2 must be 'text 3'.")

	_, ok = resource.Values["strings"].AsString()
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

	var string1 string
	string1, ok = testStrings[0].AsString()
	a.True(ok, "element 0 must be a string.")
	a.Equal("text 1", string1, "element 0 must be 'text 1'.")

	var string2 string
	string2, ok = testStrings[1].AsString()
	a.True(ok, "element 1 must be a string.")
	a.Equal("text 2", string2, "element 1 must be 'text 2'.")

	var string3 string
	string3, ok = testStrings[2].AsString()
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

	var intValue1 string
	intValue1, ok = map1["intValue"].AsString()
	a.True(ok, "'intValue' must be found in first testMap.")
	a.Equal("43", intValue1, "'intValue' must be '43'.")

	var stringValue1 string
	stringValue1, ok = map1["stringValue"].AsString()
	a.True(ok, "'stringValue' must be found in first testMap.")
	a.Equal("test 1", stringValue1, "'stringValue' must be 'test 1'.")

	var map2 map[string]ResourceData
	map2, ok = testMaps[1].AsMap()
	a.True(ok, "element 1 must be a map.")

	var intValue2 string
	intValue2, ok = map2["intValue"].AsString()
	a.True(ok, "'intValue' must be found in second testMap.")
	a.Equal("367", intValue2, "'intValue' must be '367'.")

	var stringValue2 string
	stringValue2, ok = map2["stringValue"].AsString()
	a.True(ok, "'stringValue' must be found in second testMap.")
	a.Equal("test 2", stringValue2, "'stringValue' must be 'test 2'.")
}
