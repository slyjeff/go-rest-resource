package encoding

import (
	"github.com/slyjeff/rest-resource/resource"
	"github.com/slyjeff/rest-resource/resource/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MarshalXmlMustEncodeAllProperties(t *testing.T) {
	//arrange
	i := newTestItem1()

	var r resource.Resource
	r.MapAllDataFrom(i)

	//act
	x, err := MarshalXml(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><IsAvailable>true</IsAvailable><Name>widget</Name><Price>45.2531</Price><Quantity>15</Quantity></resource>"
	a.Equal(expectedXml, string(x))
}

func Test_MarshalXmlMustEncodeFormattedData(t *testing.T) {
	//arrange
	i := newTestItem1()

	var r resource.Resource
	r.MapDataFrom(i).
		Map("Price", option.Format("%.02f"))

	//act
	x, err := MarshalXml(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><Price>45.25</Price></resource>"
	a.Equal(expectedXml, string(x))
}

func Test_MarshalXmlMustEncodeChildStructs(t *testing.T) {
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
	x, err := MarshalXml(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><Item><IsAvailable>true</IsAvailable><Name>widget</Name><Price>45.2531</Price><Quantity>15</Quantity></Item><Total>45.25</Total></resource>"
	a.Equal(expectedXml, string(x))
}

func Test_MarshalXmlMustEncodeChildSlice(t *testing.T) {
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
	x, err := MarshalXml(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedXml := "<resource><Items><Value><IsAvailable>true</IsAvailable><Name>widget</Name><Price>45.2531</Price><Quantity>15</Quantity></Value><Value><IsAvailable>false</IsAvailable><Name>thingy</Name><Price>13.84</Price><Quantity>7</Quantity></Value></Items><Total>45.25</Total></resource>"
	a.Equal(expectedXml, string(x))
}
