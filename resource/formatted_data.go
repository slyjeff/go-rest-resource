package resource

import (
	"fmt"
)

type FormattedData struct {
	Value  interface{}
	format string
}

func (fd FormattedData) FormattedString() string {
	return fmt.Sprintf(fd.format, fd.Value)
}
