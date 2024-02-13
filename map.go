package GoRestResource

func (r *Resource) MapAllDataFrom(source interface{}) *Resource {
	return r.MapDataFrom(source).MapAll().EndMap()
}

func (r *Resource) MapDataFrom(source interface{}) *ConfigureMap {
	cm := ConfigureMap{newMapFromResource(r), source}
	return &cm
}

func (r *Resource) MapAllDataFromSlice(fieldName string, source []interface{}) *Resource {
	return r.MapFromSlice(fieldName, source).MapAll().EndMap()
}

func (r *Resource) MapFromSlice(fieldName string, source []interface{}) *ConfigureSliceMapFromResource {
	fieldName = makeCamelCase(fieldName)

	slice := make([]ResourceMap, len(source))
	for i := range slice {
		slice[i] = ResourceMap{make(map[string]interface{})}
	}

	if r.Values == nil {
		r.Values = make(map[string]interface{})
	}
	r.Values[fieldName] = slice

	cm := ConfigureSliceMapFromResource{newMapFromResource(r), configureSlice{slice, source}}
	return &cm
}
