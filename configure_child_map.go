package GoRestResource

type ChildMapper interface {
	Map(fieldName string) ChildMapper
	MapWithOptions(fieldName string, mapOptions MapOptions) ChildMapper
	MapChild(fieldName string) ChildMapper
	MapAll() ChildMapper
	Exclude(fieldName string) ChildMapper
	EndMap() *Resource
}
