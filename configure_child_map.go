package GoRestResource

type ConfigureChildMap interface {
	Map(fieldName string) ConfigureChildMap
	MapWithOptions(fieldName string, mapOptions MapOptions) ConfigureChildMap
	MapAll() ConfigureChildMap
	EndMap() *ConfigureMap
}

type ConfigureChildOfResourceMap interface {
	Map(fieldName string) ConfigureChildOfResourceMap
	MapWithOptions(fieldName string, mapOptions MapOptions) ConfigureChildOfResourceMap
	MapAll() ConfigureChildOfResourceMap
	EndMap() *ConfigureResourceMap
}
