package resource

import (
	"github.com/slyjeff/rest-resource/resource/option"
)

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Link(name string, href string, linkOptions ...option.Option) *Resource {
	link := Link{Href: href, Verb: "GET", IsTemplated: false, Parameters: make([]LinkParameter, 0)}

	if verb, ok := option.FindVerbOption(linkOptions); ok {
		link.Verb = verb
	}

	link.IsTemplated = option.FindTemplatedOption(linkOptions)

	r.addLink(name, link)

	return r
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) LinkWithParameters(name string, href string, linkOptions ...option.Option) ConfigureLinkParameters {
	r.Link(name, href, linkOptions...)

	return ConfigureLinkParameters{r, r.Links[name]}
}

type ConfigureLinkParameters struct {
	resource *Resource
	link     *Link
}

func (clp ConfigureLinkParameters) Parameter(name string, parameterOptions ...option.Option) ConfigureLinkParameters {
	parameter := LinkParameter{Name: name}

	if defaultValue, ok := option.FindDefaultOption(parameterOptions); ok {
		parameter.DefaultValue = defaultValue
	}

	clp.link.Parameters = append(clp.link.Parameters, parameter)

	return clp
}

func (clp ConfigureLinkParameters) EndMap() *Resource {
	return clp.resource
}
