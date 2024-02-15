package json

import (
	"github.com/slyjeff/rest-resource/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type item struct {
	Name        string
	Quantity    int
	IsAvailable bool
	Price       float64
}

func newTestItem() item {
	return item{"widget", 15, true, 45.2531}
}

func Test_MarshalingMustEncodeAllProperties(t *testing.T) {
	//arrange
	i := newTestItem()

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
