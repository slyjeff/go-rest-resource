package resource

import (
	"github.com/slyjeff/rest-resource/option"
	"net/http"
)

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Uri(href string) *Resource {
	r.Link("self", href).
		Schema(r.Schema).
		ResponseCodes(http.StatusOK, http.StatusNotFound, http.StatusInternalServerError)

	return r
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Link(name string, href string, linkOptions ...option.Option) ConfigureLink {
	link := newLink(href)

	if verb, ok := option.FindVerbOption(linkOptions); ok {
		link.Verb = verb
	}

	link.IsTemplated = option.FindTemplatedOption(linkOptions)

	r.addLink(name, link)

	return ConfigureLink{r, r.Links[name]}
}

type ConfigureLink struct {
	resource *Resource
	link     *Link
}

func (cl ConfigureLink) Parameter(name string, parameterOptions ...option.Option) ConfigureLink {
	parameter := LinkParameter{Name: name, DefaultValue: "", ListOfValues: ""}

	if defaultValue, ok := option.FindDefaultOption(parameterOptions); ok {
		parameter.DefaultValue = defaultValue
	}

	if listOfValues, ok := option.FindListOfValuesOption(parameterOptions); ok {
		parameter.ListOfValues = listOfValues
	}

	cl.link.Parameters = append(cl.link.Parameters, parameter)

	return cl
}

func (cl ConfigureLink) Schema(schema string) ConfigureLink {
	cl.link.Schema = schema
	return cl
}

func (cl ConfigureLink) ResponseCodes(statuses ...int) ConfigureLink {
	cl.link.ResponseCodes = statuses
	return cl
}
