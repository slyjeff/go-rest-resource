package GoRestResource

import (
	"reflect"
)

func getValueByName(source interface{}, fieldName string) interface{} {
	if source == nil {
		return nil
	}

	value := reflect.ValueOf(source).FieldByName(fieldName)
	if !value.IsValid() {
		return nil
	}

	return value.Interface() //, true
}
