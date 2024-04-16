package resource

type Resource struct {
	Schema   string
	Values   MappedData
	Links    LinkData
	Embedded EmbeddedResources
}

func NewResource(schema ...string) Resource {
	s := ""
	if len(schema) > 0 {
		s = schema[0]
	}

	r := Resource{
		s,
		make(map[string]interface{}),
		make(map[string]*Link),
		make(EmbeddedResources)}

	return r
}

type Link struct {
	Href          string
	Verb          string
	IsTemplated   bool
	Parameters    []LinkParameter
	Schema        string
	ResponseCodes []int
}

func newLink(href string) Link {
	return Link{href, "GET", false, make([]LinkParameter, 0), "", make([]int, 0)}
}

type LinkParameter struct {
	Name         string
	DefaultValue string
	ListOfValues string
}

type ResponseCode struct {
	Status      int
	Description string
	Schema      string
}

type MappedData map[string]interface{}

type LinkData map[string]*Link

type EmbeddedResources map[string]interface{}

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
func (r *Resource) addEmbeddedResource(name string, resource Resource) {
	if r.Embedded == nil {
		r.Embedded = make(EmbeddedResources)
	}
	r.Embedded[name] = resource
}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) addEmbeddedResources(name string, resources []Resource) {
	if r.Embedded == nil {
		r.Embedded = make(EmbeddedResources)
	}
	r.Embedded[name] = resources
}
