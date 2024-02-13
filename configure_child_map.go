package GoRestResource

type ConfigureChildMap interface {
	Map(fieldName string) ConfigureChildMap
	MapWithOptions(fieldName string, mapOptions MapOptions) ConfigureChildMap
	EndMap() *ConfigureMap
}
