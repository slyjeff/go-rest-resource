package resource

type Resource struct {
	Values MappedData
}

type MappedData map[string]interface{}

//goland:noinspection GoMixedReceiverTypes
func (r *Resource) addData(fieldName string, value interface{}) {
	if r.Values == nil {
		r.Values = make(map[string]interface{})
	}
	r.Values[fieldName] = value
}
