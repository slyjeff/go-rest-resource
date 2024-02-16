package xml

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
	x, err := Marshal(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><IsAvailable>true</IsAvailable><Name>widget</Name><Price>45.2531</Price><Quantity>15</Quantity></resource>"
	a.Equal(expectedXml, string(x))
}

func Test_MarshalMustEncodeFormattedData(t *testing.T) {
	//arrange
	i := newTestItem1()

	var r resource.Resource
	r.MapDataFrom(i).
		Map("Price", mapping.Format("%.02f"))

	//act
	x, err := Marshal(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><Price>45.25</Price></resource>"
	a.Equal(expectedXml, string(x))
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
	x, err := Marshal(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><Item><IsAvailable>true</IsAvailable><Name>widget</Name><Price>45.2531</Price><Quantity>15</Quantity></Item><Total>45.25</Total></resource>"
	a.Equal(expectedXml, string(x))
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
	x, err := Marshal(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><Items><Value><IsAvailable>true</IsAvailable><Name>widget</Name><Price>45.2531</Price><Quantity>15</Quantity></Value><Value><IsAvailable>false</IsAvailable><Name>thingy</Name><Price>13.84</Price><Quantity>7</Quantity></Value></Items><Total>45.25</Total></resource>"
	a.Equal(expectedXml, string(x))
}
