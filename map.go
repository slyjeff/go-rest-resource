package GoRestResource

import (
	"reflect"
	"slices"
)

type copyPair struct {
	sourceItems      []interface{}
	destinationItems []*ResourceData
}

func newCopyPair(sourceItems []interface{}, destinationItems []*ResourceData) copyPair {
	return copyPair{sourceItems, destinationItems}
}

func newSingleCopyPair(sourceItem interface{}, destinationItem *ResourceData) copyPair {
	return newCopyPair([]interface{}{sourceItem}, []*ResourceData{destinationItem})
}

type ConfigureMap struct {
	resource       *Resource
	copyPairs      []copyPair
	excludedFields []string
}

func newConfigureMap(r *Resource, copyPairs ...copyPair) ConfigureMap {
	return ConfigureMap{r, copyPairs, make([]string, 0)}
}

func (r *Resource) MapChild(fieldName string, source interface{}) ChildMapper {
	sourceItems, ok := source.([]interface{})
	if ok {
		return mapChildSlice(r, &r.ResourceData, fieldName, sourceItems)
	}

	return mapChildStruct(r, &r.ResourceData, fieldName, source)
}

func (r *Resource) MapAllDataFrom(source interface{}) *Resource {
	return r.MapDataFrom(source).MapAll().EndMap()
}

func (r *Resource) MapDataFrom(source interface{}) ChildMapper {
	configuration := newConfigureMap(r, newSingleCopyPair(source, &r.ResourceData))
	return &configuration
}

func (cm *ConfigureMap) Map(fieldName string) ChildMapper {
	cm.MapWithOptions(fieldName, MapOptions{})

	return cm
}

func (cm *ConfigureMap) MapWithOptions(fieldName string, mapOptions MapOptions) ChildMapper {
	var name string
	if mapOptions.Name == "" {
		name = fieldName
	} else {
		name = mapOptions.Name
	}

	for _, copyPair := range cm.copyPairs {
		for i, v := range copyPair.sourceItems {
			resourceData := copyPair.destinationItems[i]
			if _, ok := resourceData.Values[fieldName]; ok {
				continue
			}

			value := getValueByName(v, fieldName)
			value = createResourceData(value)
			if mapOptions.FormatCallback != nil {
				value = FormattedData{value, mapOptions.FormatCallback}
			}

			resourceData.AddData(name, value)
		}
	}

	return cm
}

func (cm *ConfigureMap) MapChild(fieldName string) ChildMapper {
	copyPairs := make([]copyPair, 0)

	for _, cp := range cm.copyPairs {
		for i, v := range cp.sourceItems {
			rd := cp.destinationItems[i]
			source := getValueByName(v, fieldName)
			sourceItems, ok := source.([]interface{})
			if ok {
				destinationItems := make([]ResourceData, len(sourceItems))
				for i := range sourceItems {
					destinationItems[i] = ResourceData{make(map[string]interface{})}
				}

				rd.AddData(fieldName, destinationItems)

				destinationPointers := make([]*ResourceData, len(destinationItems))
				for i := range destinationItems {
					destinationPointers[i] = &destinationItems[i]
				}

				copyPairs = append(copyPairs, newCopyPair(sourceItems, destinationPointers))
				continue
			}

			destinationItem := ResourceData{make(map[string]interface{})}
			rd.AddData(fieldName, destinationItem)
			copyPairs = append(copyPairs, newSingleCopyPair(source, &destinationItem))
		}
	}

	newConfigureMap := newConfigureMap(cm.resource, copyPairs...)
	return &newConfigureMap
}

func mapChildSlice(r *Resource, rd *ResourceData, fieldName string, sourceItems []interface{}) ChildMapper {
	destinationItems := make([]ResourceData, len(sourceItems))
	for i := range sourceItems {
		destinationItems[i] = ResourceData{make(map[string]interface{})}
	}

	rd.AddData(fieldName, destinationItems)

	destinationPointers := make([]*ResourceData, len(destinationItems))
	for i := range destinationItems {
		destinationPointers[i] = &destinationItems[i]
	}

	copyPairs := []copyPair{{sourceItems, destinationPointers}}

	csm := newConfigureMap(r, copyPairs...)
	return &csm
}

func mapChildStruct(r *Resource, rd *ResourceData, fieldName string, source interface{}) ChildMapper {
	childResourceData := ResourceData{Values: make(map[string]interface{})}
	rd.AddData(fieldName, childResourceData)

	cm := newConfigureMap(r, newSingleCopyPair(source, &childResourceData))
	return &cm
}

func (cm *ConfigureMap) MapAll() ChildMapper {
	if len(cm.copyPairs[0].sourceItems) == 0 {
		return cm
	}

	firstItem := cm.copyPairs[0].sourceItems[0]

	t := reflect.TypeOf(firstItem)

	for i := 0; i < t.NumField(); i++ {
		var fieldName = t.Field(i).Name
		if slices.Contains(cm.excludedFields, fieldName) {
			continue
		}

		cm.MapWithOptions(fieldName, MapOptions{})
	}
	return cm
}

func (cm *ConfigureMap) Exclude(fieldName string) ChildMapper {
	cm.excludedFields = append(cm.excludedFields, fieldName)

	for _, copyPair := range cm.copyPairs {
		for i, v := range copyPair.sourceItems {
			resourceData := copyPair.destinationItems[i]

			value := getValueByName(v, fieldName)
			value = createResourceData(value)

			delete(resourceData.Values, fieldName)
		}
	}

	return cm
}

func (cm *ConfigureMap) EndMap() *Resource {
	return cm.resource
}
