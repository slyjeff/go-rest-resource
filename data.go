package restresource

import (
	"fmt"
	"reflect"
)

func (r *Resource) Data(name string, value interface{}) *Resource {
	r.addToResourceMap(name, value)
	return r
}

func (rm *resourceMap) addToResourceMap(name string, value interface{}) {
	name = makeCamelCase(name)

	if rm.Values == nil {
		rm.Values = make(map[string]ResourceData)
	}

	rm.Values[name] = createResourceData(value)
}

func createResourceData(value interface{}) ResourceData {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Struct:
		return createResourceMap(value)
	case reflect.Slice, reflect.Array:
		return createResourceSlice(value)
	default:
		return createResourceValue(value)
	}
}

func createResourceValue(value interface{}) *resourceValue {
	return &resourceValue{fmt.Sprint(value)}
}

func createResourceSlice(value interface{}) *resourceSlice {
	rs := resourceSlice{
		[]ResourceData{},
	}

	v := reflect.ValueOf(value)
	l := v.Len()
	if l == 0 {
		return &rs
	}

	values := make([]ResourceData, l)

	for i := 0; i < l; i++ {
		values[i] = createResourceData(v.Index(i).Interface())
	}

	rs.Values = values

	return &rs
}

func createResourceMap(value interface{}) *resourceMap {
	var rm resourceMap

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		rm.addToResourceMap(t.Field(i).Name, v.Field(i).Interface())
	}

	return &rm
}
