package GoRestResource

type Resource struct {
	ResourceData
}

type ResourceData struct {
	Values map[string]interface{}
}

func (rd *ResourceData) AddData(fieldName string, value interface{}) {
	if rd.Values == nil {
		rd.Values = make(map[string]interface{})
	}
	rd.Values[fieldName] = value
}
