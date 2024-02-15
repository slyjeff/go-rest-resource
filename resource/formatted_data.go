package resource

import (
	"fmt"
	"reflect"
)

type FormatDataCallback func(v interface{}) string

type FormattedData struct {
	Value  interface{}
	format string
}

func (fd FormattedData) FormattedString() string {
	return fmt.Sprintf(fd.format, fd.Value)
}

func (fd FormattedData) MarshalJSON() ([]byte, error) {
	if reflect.TypeOf(fd.Value).Kind() == reflect.String {
		return []byte("\"" + fd.FormattedString() + "\""), nil
	}

	return []byte(fd.FormattedString()), nil
}
