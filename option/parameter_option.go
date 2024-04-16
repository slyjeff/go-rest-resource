package option

import (
	"fmt"
	"strings"
)

func Default(name string) Option {
	return Option{"default", name}
}

func FindDefaultOption(options []Option) (string, bool) {
	return findOption(options, "default")
}

func ListOfValues[K any](values []K) Option {
	asStrings := make([]string, len(values))
	for i, value := range values {
		asStrings[i] = fmt.Sprintf("%v", value)
	}

	return Option{"listOfValues", strings.Join(asStrings, ",")}
}

func FindListOfValuesOption(options []Option) (string, bool) {
	return findOption(options, "listOfValues")
}

func DataType(dataType string) Option {
	return Option{"dataType", dataType}
}

func FindDataType(options []Option) (string, bool) {
	return findOption(options, "dataType")
}
