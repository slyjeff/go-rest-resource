package restresource

import (
	"golang.org/x/exp/slices"
	"reflect"
)

type ConfigureMap struct {
	resource       *Resource
	source         interface{}
	excludedFields []string
}

func (r *Resource) MapAllDataFrom(source interface{}) *Resource {
	return r.MapDataFrom(source).MapAll().EndMap()
}

func (r *Resource) MapDataFrom(source interface{}) *ConfigureMap {
	cm := ConfigureMap{r, source, make([]string, 0)}
	return &cm
}

func (cm *ConfigureMap) Map(fieldName string) *ConfigureMap {
	v := reflect.ValueOf(cm.source).FieldByName(fieldName).Interface()

	cm.resource.Data(fieldName, v)

	return cm
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
	fieldName = makeCamelCase(fieldName)
	cm.excludedFields = append(cm.excludedFields, fieldName)

	delete(cm.resource.Values, fieldName)

	return cm
}

func (cm *ConfigureMap) MapFormatted(fieldName string, callback FormatDataCallback) *ConfigureMap {
	v := reflect.ValueOf(cm.source).FieldByName(fieldName).Interface()
	fd := FormattedData{v, callback}
	cm.resource.Data(fieldName, fd)

	return cm
}

func (cm *ConfigureMap) EndMap() *Resource {
	return cm.resource
}
