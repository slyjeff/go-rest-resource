package json

import (
	"github.com/slyjeff/rest-resource/resource"
	"github.com/slyjeff/rest-resource/resource/mapping"
	"github.com/stretchr/testify/assert"
	"testing"
)

type item struct {
	Name        string
	Quantity    int
	IsAvailable bool
	Price       float64
}

func newTestItem1() item {
	return item{"widget", 15, true, 45.2531}
}
func newTestItem2() item {
	return item{"thingy", 7, false, 13.84}
}

func Test_MarshalMustEncodeAllProperties(t *testing.T) {
	//arrange
	i := newTestItem1()

	var r resource.Resource
	r.MapAllDataFrom(i)

	//act
	json, ok := Marshal(r)

	//assert
	a := assert.New(t)
	a.True(ok)
	expectedJson := `{"IsAvailable":true,"Name":"widget","Price":45.2531,"Quantity":15}`
	a.Equal(expectedJson, string(json), "json not created properly.")
}

func Test_MarshalMustEncodeFormattedData(t *testing.T) {
	//arrange
	i := newTestItem1()

	var r resource.Resource
	r.MapDataFrom(i).
		Map("Price", mapping.Format("%.02f"))

	//act
	json, ok := Marshal(r)

	//assert
	a := assert.New(t)
	a.True(ok)
	expectedJson := `{"Price":45.25}`
	a.Equal(expectedJson, string(json), "json not created properly.")
}

func Test_MarshalMustEncodeChildStructs(t *testing.T) {
	//arrange
	receipt := struct {
		Total float64
		Item  item
	}{
		45.25,
		newTestItem1(),
	}

	var r resource.Resource
	r.MapAllDataFrom(receipt)

	//act
	json, ok := Marshal(r)

	//assert
	a := assert.New(t)
	a.True(ok)
	expectedJson := `{"Item":{"IsAvailable":true,"Name":"widget","Price":45.2531,"Quantity":15},"Total":45.25}`
	a.Equal(expectedJson, string(json), "json not created properly.")
}

func Test_MarshalMustEncodeChildSlice(t *testing.T) {
	//arrange
	receipt := struct {
		Total float64
		Items []item
	}{
		45.25,
		[]item{newTestItem1(), newTestItem2()},
	}

	var r resource.Resource
	r.MapAllDataFrom(receipt)

	//act
	json, ok := Marshal(r)

	//assert
	a := assert.New(t)
	a.True(ok)
	expectedJson := `{"Items":[{"IsAvailable":true,"Name":"widget","Price":45.2531,"Quantity":15},{"IsAvailable":false,"Name":"thingy","Price":13.84,"Quantity":7}],"Total":45.25}`
	a.Equal(expectedJson, string(json), "json not created properly.")
}
