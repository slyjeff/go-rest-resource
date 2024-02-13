package GoRestResource

import (
	"reflect"
	"slices"
)

type configureSlice struct {
	slice  []ResourceMap
	source []interface{}
}

func (configuration configureSlice) mapWithOptions(fieldName string, mapOptions MapOptions) {
	if len(configuration.source) == 0 {
		return
	}

	var name string
	if mapOptions.Name == "" {
		name = makeCamelCase(fieldName)
	} else {
		name = makeCamelCase(mapOptions.Name)
	}

	for i, v := range configuration.source {
		m := configuration.slice[i]

		value := getValueByName(v, fieldName)
		if mapOptions.FormatCallback != nil {
			value = FormattedData{value, mapOptions.FormatCallback}
		}

		m.Values[name] = value
	}
}

type ConfigureSliceMapFromResource struct {
	MapFromResource MapFromResource
	configureSlice
}

type ConfigureChildSliceMap struct {
	configureResourceMap
	configureSlice
	parent *ConfigureMap
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
	if len(configuration.source) == 0 {
		return configuration
	}

	firstItem := configuration.source[0]

	t := reflect.TypeOf(firstItem)

	for i := 0; i < t.NumField(); i++ {
		var fieldName = makeCamelCase(t.Field(i).Name)
		if slices.Contains(configuration.MapFromResource.excludedFields, fieldName) {
			continue
		}

		configuration.Map(t.Field(i).Name)
	}

	return configuration
}

func (configuration *ConfigureSliceMapFromResource) Exclude(fieldName string) *ConfigureSliceMapFromResource {
	configuration.MapFromResource.excludeField(fieldName)

	return configuration
}

func (configuration *ConfigureSliceMapFromResource) EndMap() *Resource {
	return configuration.MapFromResource.resource
}

func (configuration *ConfigureChildSliceMap) Map(fieldName string) ConfigureChildMap {
	return configuration.MapWithOptions(fieldName, MapOptions{})
}

func (configuration *ConfigureChildSliceMap) MapWithOptions(fieldName string, mapOptions MapOptions) ConfigureChildMap {
	configuration.configureSlice.mapWithOptions(fieldName, mapOptions)
	return configuration
}

func (configuration *ConfigureChildSliceMap) EndMap() *ConfigureMap {
	return configuration.parent
}
