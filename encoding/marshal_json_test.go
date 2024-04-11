package encoding

import (
	"github.com/slyjeff/rest-resource/resource"
	"github.com/slyjeff/rest-resource/resource/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MarshalJsonMustEncodeAllProperties(t *testing.T) {
	//arrange
	i := newTestItem1()

	var r resource.Resource
	r.MapAllDataFrom(i)

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"IsAvailable":true,"Name":"widget","Price":45.2531,"Quantity":15}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustEncodeFormattedData(t *testing.T) {
	//arrange
	i := newTestItem1()

	var r resource.Resource
	r.MapDataFrom(i).
		Map("Price", option.Format("%.02f"))

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"Price":45.25}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustEncodeChildStructs(t *testing.T) {
	//arrange
	receipt := struct {
		Total float64
		Item  testItem
	}{
		45.25,
		newTestItem1(),
	}

	var r resource.Resource
	r.MapAllDataFrom(receipt)

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"Item":{"IsAvailable":true,"Name":"widget","Price":45.2531,"Quantity":15},"Total":45.25}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustEncodeChildSlice(t *testing.T) {
	//arrange
	receipt := struct {
		Total float64
		Items []testItem
	}{
		45.25,
		[]testItem{newTestItem1(), newTestItem2()},
	}

	var r resource.Resource
	r.MapAllDataFrom(receipt)

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"Items":[{"IsAvailable":true,"Name":"widget","Price":45.2531,"Quantity":15},{"IsAvailable":false,"Name":"thingy","Price":13.84,"Quantity":7}],"Total":45.25}`
	a.Equal(expectedJson, string(json))
}
