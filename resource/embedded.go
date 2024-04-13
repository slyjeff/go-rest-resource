package resource

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) Embed(name string, resources ...Resource) *Resource {
	r.addEmbeddedResources(name, resources)

	return r
}
