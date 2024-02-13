package restresource

import (
	"golang.org/x/exp/slices"
	"reflect"
)

type MapFromResource struct {
	resource       *Resource
	excludedFields []string
}

func NewMapFromResource(r *Resource) MapFromResource {
	return MapFromResource{r, make([]string, 0)}
}

func (mfr *MapFromResource) excludeField(fieldName string) {
	fieldName = makeCamelCase(fieldName)
	mfr.excludedFields = append(mfr.excludedFields, fieldName)

	delete(mfr.resource.Values, fieldName)
}

func (mfr *MapFromResource) EndMap() *Resource {
	return mfr.resource
}

type ConfigureMap struct {
	MapFromResource
	source interface{}
}

func (r *Resource) MapAllDataFrom(source interface{}) *Resource {
	return r.MapDataFrom(source).MapAll().EndMap()
}

func (r *Resource) MapDataFrom(source interface{}) *ConfigureMap {
	cm := ConfigureMap{NewMapFromResource(r), source}
	return &cm
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

type MapOptions struct {
	Name           string
	FormatCallback FormatDataCallback
}

type ConfigureSliceMap struct {
	MapFromResource
	slice  []interface{}
	source []interface{}
}

func (r *Resource) MapAllDataFromSlice(fieldName string, source []interface{}) *Resource {
	return r.MapFromSlice(fieldName, source).MapAll().EndMap()
}

func (r *Resource) MapFromSlice(fieldName string, source []interface{}) *ConfigureSliceMap {
	fieldName = makeCamelCase(fieldName)

	slice := make([]interface{}, len(source))
	for i := range slice {
		slice[i] = ResourceMap{make(map[string]interface{})}
	}

	if r.Values == nil {
		r.Values = make(map[string]interface{})
	}
	r.Values[fieldName] = slice

	cm := ConfigureSliceMap{NewMapFromResource(r), slice, source}
	return &cm
}

func (csm *ConfigureSliceMap) Map(fieldName string) *ConfigureSliceMap {
	csm.MapWithOptions(fieldName, MapOptions{})

	return csm
}

func (csm *ConfigureSliceMap) MapWithOptions(fieldName string, mapOptions MapOptions) *ConfigureSliceMap {
	if len(csm.source) == 0 {
		return csm
	}

	var name string
	if mapOptions.Name == "" {
		name = makeCamelCase(fieldName)
	} else {
		name = makeCamelCase(mapOptions.Name)
	}

	for i, v := range csm.source {
		m, ok := csm.slice[i].(ResourceMap)
		if !ok {
			continue
		}

		value := getValueByName(v, fieldName)
		if mapOptions.FormatCallback != nil {
			value = FormattedData{value, mapOptions.FormatCallback}
		}

		m.Values[name] = value
	}

	return csm
}

func (csm *ConfigureSliceMap) MapAll() *ConfigureSliceMap {
	if len(csm.source) == 0 {
		return csm
	}

	firstItem := csm.source[0]

	t := reflect.TypeOf(firstItem)

	for i := 0; i < t.NumField(); i++ {
		var fieldName = makeCamelCase(t.Field(i).Name)
		if slices.Contains(csm.excludedFields, fieldName) {
			continue
		}

		csm.Map(t.Field(i).Name)
	}

	return csm
}

func (csm *ConfigureSliceMap) Exclude(fieldName string) *ConfigureSliceMap {
	csm.excludeField(fieldName)

	return csm
}
