package restresource

import "reflect"

type ConfigureMap struct {
	resource *Resource
	source   interface{}
}

func (r *Resource) MapDataFrom(source interface{}) *ConfigureMap {
	cm := ConfigureMap{r, source}
	return &cm
}

func (cm *ConfigureMap) Map(fieldName string) *ConfigureMap {
	v := reflect.ValueOf(cm.source).FieldByName(fieldName).Interface()

	cm.resource.Data(fieldName, v)

	return cm
}

func (cm *ConfigureMap) EndMap() *Resource {
	return cm.resource
}
