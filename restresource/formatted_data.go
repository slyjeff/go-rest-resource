package restresource

type FormatDataCallback func(v interface{}) string

type FormattedData struct {
	Value      interface{}
	formatData FormatDataCallback
}

func (fd FormattedData) FormattedString() string {
	return fd.formatData(fd.Value)
}
