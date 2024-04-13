package resource

type Resource struct {
	Name     string
	Values   MappedData
	Links    LinkData
	Embedded EmbeddedResources
}

type Link struct {
	Href        string
	Verb        string
	IsTemplated bool
	Parameters  []LinkParameter
}

type LinkParameter struct {
	Name         string
	DefaultValue string
	ListOfValues string
}

func NewResource(name ...string) Resource {
	n := ""
	if len(name) > 0 {
		n = name[0]
	}

	r := Resource{n, make(map[string]interface{}), make(map[string]*Link), make(EmbeddedResources)}
	return r
}

type MappedData map[string]interface{}

type LinkData map[string]*Link

type EmbeddedResources map[string][]Resource

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) addData(fieldName string, value interface{}) {
	if r.Values == nil {
		r.Values = make(map[string]interface{})
	}
	r.Values[fieldName] = value
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) addLink(name string, link Link) {
	if r.Links == nil {
		r.Links = make(map[string]*Link)
	}
	r.Links[name] = &link
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) addEmbeddedResources(name string, resources []Resource) {
	if r.Embedded == nil {
		r.Embedded = make(EmbeddedResources)
	}
	r.Embedded[name] = resources
}
