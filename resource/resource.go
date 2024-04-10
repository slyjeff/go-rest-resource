package resource

type Resource struct {
	Values MappedData
	Links  map[string]*Link
}

type Link struct {
	Href       string
	Verb       string
	Parameters []LinkParameter
}

type LinkParameter struct {
	Name string
}

func NewResource() Resource {
	r := Resource{make(map[string]interface{}), make(map[string]*Link)}
	return r
}

type MappedData map[string]interface{}

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
