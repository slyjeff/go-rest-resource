package resource

import (
	"github.com/slyjeff/rest-resource/resource/option"
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
	value, ok := resource.Values["message"]
	a.True(ok, "'message' must exist")
	a.Equal(message, value, "'message' value must be 'TestMessage'")
}

func Test_DataMustStoreInt(t *testing.T) {
	//arrange
	number := 42
	var resource Resource

	//act
	resource.Data("number", number)

	//assert
	a := assert.New(t)
	value, ok := resource.Values["number"]
	a.True(ok, "'number' must exist")
	a.Equal(42, value, "'number' value must be '42'.")
}

func Test_FormattedDataAddValueAndFormattingInformation(t *testing.T) {
	//arrange
	number := 4234.3982
	var resource Resource

	//act
	resource.Data("number", number, option.Format("%.02f"))

	//assert
	a := assert.New(t)
	value, ok := resource.Values["number"]
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

	v1, ok := resource.Values["value1"]
	a.True(ok, "'value1' must exist")
	a.Equal(37, v1, "'value1' value must be '37'.")

	var v2 interface{}
	v2, ok = resource.Values["value2"]
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
	md, ok := resource.Values["testStruct"].(MappedData)
	a.True(ok, "'testStruct' must be found in values.")

	var intValue interface{}
	intValue, ok = md["IntValue"]

	a.True(ok, "'IntValue' must be in 'testStruct'.")
	a.Equal(982, intValue, "'IntValue' value must be '982'.")

	var stringValue interface{}
	stringValue, ok = md["StringValue"]

	a.True(ok, "'StringValue' must be in 'testStruct'.")
	a.Equal("Some test text.", stringValue, "'StringValue' value must be 'Some text'.")
}

func Test_DataMustTransformStructPointerToMap(t *testing.T) {
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
	resource.Data("testStruct", &testStruct)

	//assert
	a := assert.New(t)
	md, ok := resource.Values["testStruct"].(MappedData)
	a.True(ok, "'testStruct' must be found in values.")

	var intValue interface{}
	intValue, ok = md["IntValue"]

	a.True(ok, "'IntValue' must be in 'testStruct'.")
	a.Equal(982, intValue, "'IntValue' value must be '982'.")

	var stringValue interface{}
	stringValue, ok = md["StringValue"]

	a.True(ok, "'StringValue' must be in 'testStruct'.")
	a.Equal("Some test text.", stringValue, "'StringValue' value must be 'Some text'.")
}

func Test_DataMustGracefullyHandleNilPointers(t *testing.T) {
	//arrange
	type testStruct struct {
		IntValue    int
		StringValue string
	}

	var testValue *testStruct

	var resource Resource

	//act
	resource.Data("testValue", testValue)

	//assert
	a := assert.New(t)
	md, ok := resource.Values["testValue"].(MappedData)
	a.True(ok, "'testValue' must be found in values.")

	a.Equal(0, len(md), "'testValue' be an empty mapdata.")
}

func Test_DataMustNotMapPrivateFields(t *testing.T) {
	//arrange
	testStruct := struct {
		IntValue    int
		stringValue string
	}{
		IntValue:    982,
		stringValue: "Some test text.",
	}

	var resource Resource

	//act
	resource.Data("testStruct", testStruct)

	//assert
	a := assert.New(t)
	md, ok := resource.Values["testStruct"].(MappedData)
	a.True(ok, "'testStruct' must be found in values.")

	_, ok = md["stringValue"]
	a.False(ok, "'stringValue' must be in 'testStruct'.")
}

func Test_DataMustAddSliceToResource(t *testing.T) {
	//arrange
	strings := []string{"text 1", "text 2", "text 3"}
	var resource Resource

	//act
	resource.Data("strings", strings)

	//assert
	a := assert.New(t)
	s, ok := resource.Values["strings"].([]interface{})
	a.True(ok, "'strings' must be found in values.")

	a.Equal("text 1", s[0], "element 0 must be 'text 1'.")
	a.Equal("text 2", s[1], "element 1 must be 'text 2'.")
	a.Equal("text 3", s[2], "element 2 must be 'text 3'.")
}

func Test_DataMustAddArrayToResource(t *testing.T) {
	//arrange
	strings := [...]string{"text 1", "text 2", "text 3"}
	var resource Resource

	//act
	resource.Data("strings", strings)

	//assert
	a := assert.New(t)
	s, ok := resource.Values["strings"].([]interface{})
	a.True(ok, "'strings' must be found in values.")

	a.Equal("text 1", s[0], "element 0 must be 'text 1'.")
	a.Equal("text 2", s[1], "element 1 must be 'text 2'.")
	a.Equal("text 3", s[2], "element 2 must be 'text 3'.")
}

func Test_DatMustTransformStructsToResourceMapsInSlices(t *testing.T) {
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
	slice, ok := resource.Values["structs"].([]MappedData)
	a.True(ok, "'structs' must be found in values.")

	var intValue1 interface{}
	intValue1, ok = slice[0]["IntValue"]
	a.True(ok, "'IntValue' must be found in first testMap.")
	a.Equal(43, intValue1, "'IntValue' must be '43'.")

	var stringValue1 interface{}
	stringValue1, ok = slice[0]["StringValue"]
	a.True(ok, "'StringValue' must be found in first testMap.")
	a.Equal("test 1", stringValue1, "'StringValue' must be 'test 1'.")

	var intValue2 interface{}
	intValue2, ok = slice[1]["IntValue"]
	a.True(ok, "'IntValue' must be found in second testMap.")
	a.Equal(367, intValue2, "'IntValue' must be '367'.")

	var stringValue2 interface{}
	stringValue2, ok = slice[1]["StringValue"]
	a.True(ok, "'StringValue' must be found in second testMap.")
	a.Equal("test 2", stringValue2, "'StringValue' must be 'test 2'.")
}
