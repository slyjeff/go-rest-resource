package GoRestResource

import (
	"reflect"
	"slices"
)

type ConfigureMap struct {
	resource       *Resource
	resourceData   *ResourceData
	source         interface{}
	excludedFields []string
}

func newConfigureMap(r *Resource, rd *ResourceData, source interface{}) ConfigureMap {
	return ConfigureMap{r, rd, source, make([]string, 0)}
}

func (r *Resource) MapChild(fieldName string, source interface{}) ChildMapper {
	sourceItems, ok := source.([]interface{})
	if !ok {
		sourceItems = make([]interface{}, 0)
	}

	return mapChild(r, &r.ResourceData, fieldName, sourceItems)
}

func (r *Resource) MapAllDataFrom(source interface{}) *Resource {
	return r.MapDataFrom(source).MapAll().EndMap()
}

func (r *Resource) MapDataFrom(source interface{}) ChildMapper {
	configuration := newConfigureMap(r, &r.ResourceData, source)
	return &configuration
}

func (cm *ConfigureMap) Map(fieldName string) ChildMapper {
	cm.MapWithOptions(fieldName, MapOptions{})

	return cm
}

func (cm *ConfigureMap) MapWithOptions(fieldName string, mapOptions MapOptions) ChildMapper {
	v := getValueByName(cm.source, fieldName)
	v = createResourceData(v)

	if mapOptions.Name != "" {
		fieldName = mapOptions.Name
	}

	if mapOptions.FormatCallback != nil {
		v = FormattedData{v, mapOptions.FormatCallback}
	}

	if cm.resourceData.Values == nil {
		cm.resourceData.Values = make(map[string]interface{})
	}

	name := makeCamelCase(fieldName)
	cm.resourceData.Values[name] = v

	return cm
}

func (cm *ConfigureMap) MapChild(fieldName string) ChildMapper {
	sourceItems, ok := getValueByName(cm.source, fieldName).([]interface{})
	if !ok {
		sourceItems = make([]interface{}, 0)
	}
	return mapChild(cm.resource, cm.resourceData, fieldName, sourceItems)
}

func mapChild(r *Resource, rd *ResourceData, fieldName string, sourceItems []interface{}) ChildMapper {
	destinationItems := make([]ResourceData, len(sourceItems))
	for i := range sourceItems {
		destinationItems[i] = ResourceData{make(map[string]interface{})}
	}

	if rd.Values == nil {
		rd.Values = make(map[string]interface{})
	}

	name := makeCamelCase(fieldName)
	rd.Values[name] = destinationItems

	copyPairs := []copyPair{{sourceItems, destinationItems}}

	csm := newConfigureSliceMap(r, copyPairs)
	return &csm
}

func (cm *ConfigureMap) MapAll() ChildMapper {
	t := reflect.TypeOf(cm.source)
	v := reflect.ValueOf(cm.source)

	for i := 0; i < t.NumField(); i++ {
		fieldName := makeCamelCase(t.Field(i).Name)

		if slices.Contains(cm.excludedFields, fieldName) {
			continue
		}

		if _, ok := cm.resourceData.Values[fieldName]; ok {
			continue
		}

		if cm.resourceData.Values == nil {
			cm.resourceData.Values = make(map[string]interface{})
		}

		v := v.Field(i).Interface()
		name := makeCamelCase(fieldName)
		cm.resourceData.Values[name] = v
	}

	return cm
}

func (cm *ConfigureMap) Exclude(fieldName string) ChildMapper {
	fieldName = makeCamelCase(fieldName)
	cm.excludedFields = append(cm.excludedFields, fieldName)

	delete(cm.resourceData.Values, fieldName)

	return cm
}

func (cm *ConfigureMap) EndMap() *Resource {
	return cm.resource
}
