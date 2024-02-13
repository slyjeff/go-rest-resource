package GoRestResource

import (
	"reflect"
	"slices"
)

type ConfigureMap struct {
	MapFromResource
	source interface{}
}

func (cm *ConfigureMap) Map(fieldName string) *ConfigureMap {
	cm.MapWithOptions(fieldName, MapOptions{})

	return cm
}

func (cm *ConfigureMap) MapWithOptions(fieldName string, mapOptions MapOptions) *ConfigureMap {
	v := getValueByName(cm.source, fieldName)

	if mapOptions.Name != "" {
		fieldName = mapOptions.Name
	}

	if mapOptions.FormatCallback != nil {
		v = FormattedData{v, mapOptions.FormatCallback}
	}

	cm.resource.Data(fieldName, v)

	return cm
}

func (cm *ConfigureMap) MapChild(fieldName string) ConfigureChildMap {
	source, ok := getValueByName(cm.source, fieldName).([]interface{})
	if !ok {
		source = make([]interface{}, 0)
	}

	slice := make([]ResourceMap, len(source))
	for i := range slice {
		slice[i] = ResourceMap{make(map[string]interface{})}
	}

	if cm.resource.Values == nil {
		cm.resource.Values = make(map[string]interface{})
	}

	fieldName = makeCamelCase(fieldName)
	cm.resource.Values[fieldName] = slice

	mapFromResource := newMapFromResource(cm.resource)
	ccm := ConfigureChildSliceMap{&mapFromResource, configureSlice{slice, source}, cm}
	return &ccm
}

func (cm *ConfigureMap) MapAll() *ConfigureMap {
	t := reflect.TypeOf(cm.source)
	v := reflect.ValueOf(cm.source)

	for i := 0; i < t.NumField(); i++ {
		fieldName := makeCamelCase(t.Field(i).Name)

		if slices.Contains(cm.excludedFields, fieldName) {
			continue
		}

		if _, ok := cm.resource.Values[fieldName]; ok {
			continue
		}

		value := v.Field(i).Interface()
		cm.resource.Data(fieldName, value)
	}

	return cm
}

func (cm *ConfigureMap) Exclude(fieldName string) *ConfigureMap {
	cm.excludeField(fieldName)

	return cm
}
