package GoRestResource

import (
	"reflect"
	"slices"
)

type configureSlice struct {
	ConfigureMap
	slice []ResourceMap
}

func newConfigureSlice(rm *ResourceMap, slice []ResourceMap, source []interface{}) configureSlice {
	configureMap := newConfigureMap(rm, source)
	return configureSlice{configureMap, slice}
}

func (configuration configureSlice) mapWithOptions(fieldName string, mapOptions MapOptions) {
	sourceSlice, ok := configuration.source.([]interface{})
	if !ok {
		return
	}

	if len(sourceSlice) == 0 {
		return
	}

	var name string
	if mapOptions.Name == "" {
		name = makeCamelCase(fieldName)
	} else {
		name = makeCamelCase(mapOptions.Name)
	}

	for i, v := range sourceSlice {
		m := configuration.slice[i]

		value := getValueByName(v, fieldName)
		value = createResourceData(value)
		if mapOptions.FormatCallback != nil {
			value = FormattedData{value, mapOptions.FormatCallback}
		}

		m.Values[name] = value
	}
}

func (configuration configureSlice) mapAll() {
	sourceSlice, ok := configuration.source.([]interface{})
	if !ok {
		return
	}

	if len(sourceSlice) == 0 {
		return
	}

	firstItem := sourceSlice[0]

	t := reflect.TypeOf(firstItem)

	for i := 0; i < t.NumField(); i++ {
		var fieldName = makeCamelCase(t.Field(i).Name)
		if slices.Contains(configuration.excludedFields, fieldName) {
			continue
		}

		configuration.mapWithOptions(t.Field(i).Name, MapOptions{})
	}
}

type ConfigureSliceMapFromResource struct {
	configureSlice
	resource *Resource
}

func newConfigureSliceMapFromResource(r *Resource, slice []ResourceMap, source []interface{}) ConfigureSliceMapFromResource {
	return ConfigureSliceMapFromResource{newConfigureSlice(&r.ResourceMap, slice, source), r}
}

type ConfigureChildSliceOfResourceMap struct {
	configureSlice
	parent *ConfigureResourceMap
}

func newConfigureChildSliceOfResourceMap(parent *ConfigureResourceMap, slice []ResourceMap, source []interface{}) ConfigureChildSliceOfResourceMap {
	return ConfigureChildSliceOfResourceMap{newConfigureSlice(parent.ResourceMap, slice, source), parent}
}

func (configuration *ConfigureSliceMapFromResource) Map(fieldName string) *ConfigureSliceMapFromResource {
	configuration.MapWithOptions(fieldName, MapOptions{})

	return configuration
}

func (configuration *ConfigureSliceMapFromResource) MapWithOptions(fieldName string, mapOptions MapOptions) *ConfigureSliceMapFromResource {
	configuration.configureSlice.mapWithOptions(fieldName, mapOptions)
	return configuration
}

func (configuration *ConfigureSliceMapFromResource) MapAll() *ConfigureSliceMapFromResource {
	configuration.configureSlice.mapAll()
	return configuration
}

func (configuration *ConfigureSliceMapFromResource) Exclude(fieldName string) *ConfigureSliceMapFromResource {
	configuration.excludeField(fieldName)

	return configuration
}

func (configuration *ConfigureSliceMapFromResource) EndMap() *Resource {
	return configuration.resource
}

func (configuration *ConfigureChildSliceOfResourceMap) Map(fieldName string) ConfigureChildOfResourceMap {
	configuration.MapWithOptions(fieldName, MapOptions{})
	return configuration
}

func (configuration *ConfigureChildSliceOfResourceMap) MapWithOptions(fieldName string, mapOptions MapOptions) ConfigureChildOfResourceMap {
	configuration.configureSlice.mapWithOptions(fieldName, mapOptions)
	return configuration
}

func (configuration *ConfigureChildSliceOfResourceMap) MapAll() ConfigureChildOfResourceMap {
	configuration.configureSlice.mapAll()
	return configuration
}

func (configuration *ConfigureChildSliceOfResourceMap) EndMap() *ConfigureResourceMap {
	return configuration.parent
}
