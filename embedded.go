package resource

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) EmbedResource(name string, resource Resource) *Resource {
	r.addEmbeddedResource(name, resource)

	return r
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) EmbedResources(name string, resources []Resource) *Resource {
	r.addEmbeddedResources(name, resources)

	return r
}
