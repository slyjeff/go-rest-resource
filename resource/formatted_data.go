package resource

import (
	"fmt"
	"reflect"
)

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
