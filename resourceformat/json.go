package resourceformat

import (
	"RestResource/restresource"
	"fmt"
	"reflect"
	"strings"
)

func ToSlySoftHalJson(r restresource.Resource) string {
	var valueStrings []string
	for k, v := range r.Values {
		value, ok := v.AsValue()
		if !ok {
			continue
		}

		var sValue string
		if reflect.TypeOf(value).Kind() == reflect.String {
			sValue = fmt.Sprint("\"", value, "\"")
		} else {
			sValue = fmt.Sprint(value)
		}
		valueStrings = append(valueStrings, fmt.Sprintf("  \"%s\":%s", k, sValue))
	}

	return "{\n" + strings.Join(valueStrings, ",\n") + "\n}"
}
