package encoding

import (
	"github.com/slyjeff/rest-resource"
	"github.com/slyjeff/rest-resource/option"
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

func Test_MarshalJsonMustEncodeLinks(t *testing.T) {
	//arrange
	var r resource.Resource
	r.Link("getUsers", "/user")
	r.Link("getMessages", "/message")

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"_links":{"getMessages":{"href":"/message"},"getUsers":{"href":"/user"}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustEncodeLinksAndData(t *testing.T) {
	//arrange
	var r resource.Resource
	r.Data("message", "Hello World!")
	r.Link("getUsers", "/user")
	r.Link("getMessages", "/message")

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"message":"Hello World!","_links":{"getMessages":{"href":"/message"},"getUsers":{"href":"/user"}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustOutputVerbIfNotGet(t *testing.T) {
	//arrange
	var r resource.Resource
	r.Link("createUser", "/user", option.Verb("POST"))

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"_links":{"createUser":{"href":"/user","verb":"POST"}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustOutputTemplatedIfSet(t *testing.T) {
	//arrange
	var r resource.Resource
	r.Link("createUser", "/user", option.Templated())

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"_links":{"createUser":{"href":"/user","templated":true}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustOutputLinkParameters(t *testing.T) {
	//arrange
	var r resource.Resource
	r.Link("createUser", "/user").
		Parameter("param1").
		Parameter("param2")

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"_links":{"createUser":{"href":"/user","parameters":{"param1":{},"param2":{}}}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustOutputLinkParameterDefaultValues(t *testing.T) {
	//arrange
	var r resource.Resource
	r.Link("createUser", "/user").
		Parameter("param1", option.Default("max"))

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"_links":{"createUser":{"href":"/user","parameters":{"param1":{"default":"max"}}}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustOutputLinkParameterListOfValues(t *testing.T) {
	//arrange
	var r resource.Resource
	listOfValues := []int{1, 2, 3}
	r.Link("createUser", "/user").
		Parameter("param1", option.ListOfValues(listOfValues))

	//act
	json, err := MarshalJson(r)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"_links":{"createUser":{"href":"/user","parameters":{"param1":{"listOfValues":"1,2,3"}}}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustOutputEmbeddedResource(t *testing.T) {
	//arrange
	var parent = resource.NewResource("parent")
	parent.Data("id", 1)

	var child = resource.NewResource("child")
	child.Data("id", 2)

	parent.EmbedResource("child", child)

	//act
	json, err := MarshalJson(parent)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"id":1,"_embedded":{"child":{"id":2}}}`
	a.Equal(expectedJson, string(json))
}

func Test_MarshalJsonMustOutputEmbeddedResourceList(t *testing.T) {
	//arrange
	var parent = resource.NewResource("parent")
	parent.Data("id", 1)

	var child1 = resource.NewResource("child")
	child1.Data("id", 2)

	var child2 = resource.NewResource("child")
	child2.Data("id", 3)

	parent.EmbedResources("children", []resource.Resource{child1, child2})

	//act
	json, err := MarshalJson(parent)

	//assert
	a := assert.New(t)
	a.NoError(err)
	expectedJson := `{"id":1,"_embedded":{"children":[{"id":2},{"id":3}]}}`
	a.Equal(expectedJson, string(json))
}
