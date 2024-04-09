package resource

import (
	"github.com/slyjeff/rest-resource/resource/linkoption"
)

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Link(name string, href string, linkOptions ...linkoption.LinkOption) *Resource {
	if verb, ok := linkoption.FindVerbOption(linkOptions); ok {
		r.addLink(name, Link{Href: href, Verb: verb})
	} else {
		r.addLink(name, Link{Href: href, Verb: "GET"})
	}

	return r
}
