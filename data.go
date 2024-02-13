package GoRestResource

import (
	"reflect"
)

func (r *Resource) Data(name string, value interface{}) *Resource {
	r.addToResourceMap(name, value)
	return r
}

func (r *Resource) FormattedData(name string, value interface{}, callback FormatDataCallback) *Resource {
	fd := FormattedData{value, callback}
	r.addToResourceMap(name, fd)
	return r
}

func (rm *ResourceMap) addToResourceMap(name string, value interface{}) {
	name = makeCamelCase(name)

	if rm.Values == nil {
		rm.Values = make(map[string]interface{})
	}

	rm.Values[name] = createResourceData(value)
}

func createResourceData(value interface{}) interface{} {
	_, ok := value.(FormattedData)
	if ok {
		return value
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Struct:
		return createResourceMap(value)
	case reflect.Slice, reflect.Array:
		return createResourceSlice(value)
	default:
		return value
	}
}

func createResourceSlice(value interface{}) interface{} {
	v := reflect.ValueOf(value)
	l := v.Len()

	slice := make([]interface{}, l)

	for i := 0; i < l; i++ {
		slice[i] = createResourceData(v.Index(i).Interface())
	}

	return slice
}

func createResourceMap(value interface{}) interface{} {
	var rm ResourceMap

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		rm.addToResourceMap(t.Field(i).Name, v.Field(i).Interface())
	}

	return rm
}
