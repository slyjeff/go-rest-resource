package GoRestResource

import (
	"reflect"
	"slices"
)

type ConfigureMap struct {
	*ResourceMap
	source         interface{}
	excludedFields []string
}

func newConfigureMap(rm *ResourceMap, source interface{}) ConfigureMap {
	return ConfigureMap{rm, source, make([]string, 0)}
}

func (configuration *ConfigureMap) mapWithOptions(fieldName string, mapOptions MapOptions) {
	v := getValueByName(configuration.source, fieldName)
	v = createResourceData(v)

	if mapOptions.Name != "" {
		fieldName = mapOptions.Name
	}

	if mapOptions.FormatCallback != nil {
		v = FormattedData{v, mapOptions.FormatCallback}
	}

	if configuration.Values == nil {
		configuration.Values = make(map[string]interface{})
	}

	name := makeCamelCase(fieldName)
	configuration.Values[name] = v
}

func (configuration *ConfigureMap) mapChild(fieldName string) ([]ResourceMap, []interface{}) {
	source, ok := getValueByName(configuration.source, fieldName).([]interface{})
	if !ok {
		source = make([]interface{}, 0)
	}

	slice := make([]ResourceMap, len(source))
	for i := range slice {
		slice[i] = ResourceMap{make(map[string]interface{})}
	}

	if configuration.Values == nil {
		configuration.Values = make(map[string]interface{})
	}

	fieldName = makeCamelCase(fieldName)
	configuration.Values[fieldName] = slice

	return slice, source
}

func (configuration *ConfigureMap) excludeField(fieldName string) {
	fieldName = makeCamelCase(fieldName)
	configuration.excludedFields = append(configuration.excludedFields, fieldName)

	delete(configuration.Values, fieldName)
}

func (configuration *ConfigureMap) mapAll() {
	t := reflect.TypeOf(configuration.source)
	v := reflect.ValueOf(configuration.source)

	for i := 0; i < t.NumField(); i++ {
		fieldName := makeCamelCase(t.Field(i).Name)

		if slices.Contains(configuration.excludedFields, fieldName) {
			continue
		}

		if _, ok := configuration.Values[fieldName]; ok {
			continue
		}

		if configuration.Values == nil {
			configuration.Values = make(map[string]interface{})
		}

		v := v.Field(i).Interface()
		name := makeCamelCase(fieldName)
		configuration.Values[name] = v
	}
}

func (configuration *ConfigureMap) Map(fieldName string) *ConfigureMap {
	configuration.MapWithOptions(fieldName, MapOptions{})
	return configuration
}

func (configuration *ConfigureMap) MapWithOptions(fieldName string, mapOptions MapOptions) *ConfigureMap {
	configuration.MapWithOptions(fieldName, mapOptions)
	return configuration
}

func (configuration *ConfigureMap) MapAll() *ConfigureMap {
	configuration.mapAll()

	return configuration
}

func (configuration *ConfigureMap) Exclude(fieldName string) *ConfigureMap {
	configuration.excludeField(fieldName)

	return configuration
}
