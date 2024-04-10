package resource

import (
	"github.com/slyjeff/rest-resource/resource/linkoption"
)

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Link(name string, href string, linkOptions ...linkoption.LinkOption) *Resource {
	if verb, ok := linkoption.FindVerbOption(linkOptions); ok {
		r.addLink(name, Link{Href: href, Verb: verb, Parameters: make([]LinkParameter, 0)})
	} else {
		r.addLink(name, Link{Href: href, Verb: "GET", Parameters: make([]LinkParameter, 0)})
	}

	return r
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) LinkWithParameters(name string, href string, linkOptions ...linkoption.LinkOption) ConfigureLinkParameters {
	r.Link(name, href, linkOptions...)

	return ConfigureLinkParameters{r, r.Links[name]}
}

type ConfigureLinkParameters struct {
	resource *Resource
	link     *Link
}

func (clp ConfigureLinkParameters) Parameter(name string) ConfigureLinkParameters {
	clp.link.Parameters = append(clp.link.Parameters, LinkParameter{Name: name})
	return clp
}

func (clp ConfigureLinkParameters) EndMap() *Resource {
	return clp.resource
}
