package GoRestResource

type ConfigureResourceMap struct {
	ConfigureMap
	resource *Resource
}

func newConfigureResourceMap(parent *Resource, source interface{}) ConfigureResourceMap {
	return ConfigureResourceMap{newConfigureMap(&parent.ResourceMap, source), parent}
}

func (r *Resource) MapAllDataFrom(source interface{}) *Resource {
	return r.MapDataFrom(source).MapAll().EndMap()
}

func (r *Resource) MapDataFrom(source interface{}) *ConfigureResourceMap {
	configuration := newConfigureResourceMap(r, source)
	return &configuration
}

func (configuration *ConfigureResourceMap) Map(fieldName string) *ConfigureResourceMap {
	configuration.mapWithOptions(fieldName, MapOptions{})

	return configuration
}

func (configuration *ConfigureResourceMap) MapWithOptions(fieldName string, mapOptions MapOptions) *ConfigureResourceMap {
	configuration.mapWithOptions(fieldName, mapOptions)

	return configuration
}

func (configuration *ConfigureResourceMap) MapChild(fieldName string) ConfigureChildOfResourceMap {
	slice, source := configuration.mapChild(fieldName)

	childConfiguration := newConfigureChildSliceOfResourceMap(configuration, slice, source)
	return &childConfiguration
}

func (configuration *ConfigureResourceMap) MapAll() *ConfigureResourceMap {
	configuration.mapAll()
	return configuration
}

func (configuration *ConfigureResourceMap) EndMap() *Resource {
	return configuration.resource
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

	configuration := newConfigureSliceMapFromResource(r, slice, source)
	return &configuration
}
