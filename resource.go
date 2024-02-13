package GoRestResource

type Resource struct {
	ResourceMap
}

type ResourceMap struct {
	Values map[string]interface{}
}

func (rm *ResourceMap) GetResourceMap() *ResourceMap {
	return rm
}
