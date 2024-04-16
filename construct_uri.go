package resource

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func ConstructUriFromTemplate(template string, parameters ...interface{}) string {
	re := regexp.MustCompile("{[a-zA-Z0-9]*}")
	parameterNames := re.FindAllString(template, -1)

	for i, parameter := range parameters {
		if len(parameterNames) <= i {
			break
		}

		if isZeroOfUnderlyingType(parameter) {
			continue
		}

		template = strings.Replace(template, parameterNames[i], fmt.Sprintf("%v", parameter), 1)
	}
	return template
}

func isZeroOfUnderlyingType(x interface{}) bool {
	//nabbed from here: https://stackoverflow.com/questions/13901819/quick-way-to-detect-empty-values-via-reflection-in-go
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
