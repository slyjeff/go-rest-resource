package GoRestResource

import (
	"reflect"
)

func getValueByName(source interface{}, fieldName string) interface{} {
	return reflect.ValueOf(source).FieldByName(fieldName).Interface()
}
