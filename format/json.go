package format

import (
	"fmt"
	"reflect"
	"restresource"
	"strings"
)

func ToSlySoftHalJson(r GoRestResource.Resource) string {
	var valueStrings []string
	for k, v := range r.Values {
		var sValue string
		if reflect.TypeOf(v).Kind() == reflect.String {
			sValue = fmt.Sprint("\"", v, "\"")
		} else {
			sValue = fmt.Sprint(v)
		}
		valueStrings = append(valueStrings, fmt.Sprintf("  \"%s\":%s", k, sValue))
	}

	return "{\n" + strings.Join(valueStrings, ",\n") + "\n}"
}
