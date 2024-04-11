package resource

import (
	"github.com/slyjeff/rest-resource/resource/linkoption"
)

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Link(name string, href string, linkOptions ...linkoption.LinkOption) *Resource {
	link := Link{Href: href, Verb: "GET", IsTemplated: false, Parameters: make([]LinkParameter, 0)}

	if verb, ok := linkoption.FindVerbOption(linkOptions); ok {
		link.Verb = verb
	}

	link.IsTemplated = linkoption.FindTemplatedOption(linkOptions)

	r.addLink(name, link)

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
