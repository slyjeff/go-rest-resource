package restresource

import "reflect"

type MapFromResource struct {
	resource *Resource
}

func (cm *MapFromResource) EndMap() *Resource {
	return cm.resource
}

type ConfigureMap struct {
	MapFromResource
	source interface{}
}

func (r *Resource) MapDataFrom(source interface{}) *ConfigureMap {
	cm := ConfigureMap{MapFromResource{r}, source}
	return &cm
}

func (cm *ConfigureMap) Map(fieldName string) *ConfigureMap {
	v := reflect.ValueOf(cm.source).FieldByName(fieldName).Interface()

	cm.resource.Data(fieldName, v)

	return cm
}

func (cm *ConfigureMap) MapFormatted(fieldName string, callback FormatDataCallback) *ConfigureMap {
	v := reflect.ValueOf(cm.source).FieldByName(fieldName).Interface()
	fd := FormattedData{v, callback}
	cm.resource.Data(fieldName, fd)

	return cm
}

type ConfigureSliceMap struct {
	MapFromResource
	slice  []ResourceData
	source []interface{}
}

func (r *Resource) MapSliceFrom(fieldName string, source []interface{}) *ConfigureSliceMap {
	fieldName = makeCamelCase(fieldName)

	slice := make([]ResourceData, len(source))
	for i := range slice {
		slice[i] = &resourceMap{Values: map[string]ResourceData{}}
	}

	rs := resourceSlice{
		slice,
	}

	if r.Values == nil {
		r.Values = make(map[string]ResourceData)
	}
	r.Values[fieldName] = &rs

	cm := ConfigureSliceMap{MapFromResource{r}, slice, source}
	return &cm
}

func (csm *ConfigureSliceMap) Map(fieldName string) *ConfigureSliceMap {
	fieldName = makeCamelCase(fieldName)

	for i, v := range csm.source {
		m, ok := csm.slice[i].(*resourceMap)
		if !ok {
			continue
		}

		m.Values[fieldName] = createResourceValue(v)

		//if m, ok := csm.slice[i].AsMap(); !ok {
		//	m[fieldName] = createResourceValue(v)
		//}
	}

	return csm
}
