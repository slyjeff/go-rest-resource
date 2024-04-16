package resource

import (
	"github.com/slyjeff/rest-resource/option"
	"net/http"
)

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Uri(href string) *Resource {
	r.Link("self", href).
		ResponseSchema(r.Schema)

	return r
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Link(name string, href string, linkOptions ...option.Option) ConfigureLink {
	link := newLink(href)

	if verb, ok := option.FindVerbOption(linkOptions); ok {
		link.Verb = verb
		switch link.Verb {
		case "POST":
			link.ResponseCodes = []int{http.StatusCreated, http.StatusBadRequest, http.StatusInternalServerError}
		case "PUT":
			link.ResponseCodes = []int{http.StatusOK, http.StatusBadRequest, http.StatusNotFound, http.StatusInternalServerError}
		case "PATCH":
			link.ResponseCodes = []int{http.StatusOK, http.StatusBadRequest, http.StatusNotFound, http.StatusInternalServerError}
		case "DELETE":
			link.ResponseCodes = []int{http.StatusOK, http.StatusNotFound, http.StatusInternalServerError}
		}
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
	parameter := newLinkParameter(name)

	if defaultValue, ok := option.FindDefaultOption(parameterOptions); ok {
		parameter.DefaultValue = defaultValue
	}

	if listOfValues, ok := option.FindListOfValuesOption(parameterOptions); ok {
		parameter.ListOfValues = listOfValues
	}

	if dataType, ok := option.FindDataType(parameterOptions); ok {
		parameter.DataType = dataType
	}

	cl.link.Parameters = append(cl.link.Parameters, parameter)

	return cl
}

func (cl ConfigureLink) ResponseSchema(responseSchema string) ConfigureLink {
	cl.link.ResponseSchema = responseSchema
	return cl
}

func (cl ConfigureLink) ResponseCodes(statuses ...int) ConfigureLink {
	cl.link.ResponseCodes = statuses
	return cl
}
