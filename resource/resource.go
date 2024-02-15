package resource

type Resource struct {
	MappedData
}

type MappedData struct {
	Values map[string]interface{}
}

func (md *MappedData) AddData(fieldName string, value interface{}) {
	if md.Values == nil {
		md.Values = make(map[string]interface{})
	}
	md.Values[fieldName] = value
}
