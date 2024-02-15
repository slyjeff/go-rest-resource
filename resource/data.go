package resource

import (
	"reflect"
)

func (r *Resource) Data(name string, value interface{}) *Resource {
	value = createResourceData(value)
	r.addData(name, value)
	return r
}

func (r *Resource) FormattedData(name string, value interface{}, callback FormatDataCallback) *Resource {
	fd := FormattedData{value, callback}
	r.addData(name, fd)
	return r
}

func createResourceData(value interface{}) interface{} {
	if value == nil {
		return nil
	}

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

	if l == 0 {
		return make([]interface{}, 0)
	}

	firstItem := v.Index(0).Interface()
	if reflect.TypeOf(firstItem).Kind() == reflect.Struct {
		slice := make([]MappedData, l)

		for i := 0; i < l; i++ {
			slice[i] = createResourceMap(v.Index(i).Interface())
		}

		return slice
	}

	slice := make([]interface{}, l)

	for i := 0; i < l; i++ {
		slice[i] = createResourceData(v.Index(i).Interface())
	}

	return slice
}

func createResourceMap(value interface{}) MappedData {
	rm := make(MappedData)

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		v := createResourceData(v.Field(i).Interface())
		rm[t.Field(i).Name] = v
	}

	return rm
}