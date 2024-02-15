package resource

type FormatDataCallback func(v interface{}) string

type FormattedData struct {
	Value      interface{}
	formatData FormatDataCallback
}

func (fd FormattedData) FormattedString() string {
	return fd.formatData(fd.Value)
}

func (fd FormattedData) MarshalJSON() ([]byte, error) {
	return []byte(fd.FormattedString()), nil
}
