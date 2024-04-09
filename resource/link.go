package resource

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Link(name string, href string) *Resource {
	r.addLink(name, Link{Href: href, Verb: "GET"})

	return r
}
