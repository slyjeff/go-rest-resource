package GoRestResource

type MapFromResource struct {
	resource       *Resource
	excludedFields []string
}

func newMapFromResource(r *Resource) MapFromResource {
	return MapFromResource{r, make([]string, 0)}
}

func (mfr *MapFromResource) GetResourceMap() *ResourceMap {
	return &mfr.resource.ResourceMap
}

func (mfr *MapFromResource) excludeField(fieldName string) {
	fieldName = makeCamelCase(fieldName)
	mfr.excludedFields = append(mfr.excludedFields, fieldName)

	delete(mfr.GetResourceMap().Values, fieldName)
}

func (mfr *MapFromResource) EndMap() *Resource {
	return mfr.resource
}
