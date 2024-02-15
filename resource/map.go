package resource

import (
	"reflect"
	"slices"
)

type copyPair struct {
	sourceItems      []interface{}
	destinationItems []*MappedData
}

func newCopyPair(sourceItems []interface{}, destinationItems []*MappedData) copyPair {
	return copyPair{sourceItems, destinationItems}
}

func newSingleCopyPair(sourceItem interface{}, destinationItem *MappedData) copyPair {
	return newCopyPair([]interface{}{sourceItem}, []*MappedData{destinationItem})
}

type ConfigureMap struct {
	resource       *Resource
	copyPairs      []copyPair
	excludedFields []string
}

func newConfigureMap(r *Resource, copyPairs ...copyPair) ConfigureMap {
	return ConfigureMap{r, copyPairs, make([]string, 0)}
}

func (r *Resource) MapChild(fieldName string, source interface{}) *ConfigureMap {
	if r.Values == nil {
		r.Values = make(MappedData)
	}
	sourceItems, ok := source.([]interface{})
	if ok {
		return mapChildSlice(r, &r.Values, fieldName, sourceItems)
	}

	return mapChildStruct(r, &r.Values, fieldName, source)
}

func (r *Resource) MapAllDataFrom(source interface{}) *Resource {
	return r.MapDataFrom(source).MapAll().EndMap()
}

func (r *Resource) MapDataFrom(source interface{}) *ConfigureMap {
	if r.Values == nil {
		r.Values = make(MappedData)
	}
	configuration := newConfigureMap(r, newSingleCopyPair(source, &r.Values))
	return &configuration
}

func (cm *ConfigureMap) Map(fieldName string, mapOptions ...MapOption) *ConfigureMap {
	name := fieldName
	if newName, ok := FindNameOption(mapOptions); ok {
		name = newName
	}

	for _, copyPair := range cm.copyPairs {
		for i, v := range copyPair.sourceItems {
			resourceData := *copyPair.destinationItems[i]
			if _, ok := resourceData[fieldName]; ok {
				continue
			}

			value := getValueByName(v, fieldName)
			value = createResourceData(value)

			if format, ok := FindFormatOption(mapOptions); ok {
				value = FormattedData{value, format}
			}

			resourceData[name] = value
		}
	}

	return cm
}

func (cm *ConfigureMap) MapChild(fieldName string) *ConfigureMap {
	copyPairs := make([]copyPair, 0)

	for _, cp := range cm.copyPairs {
		for i, v := range cp.sourceItems {
			md := *cp.destinationItems[i]
			source := getValueByName(v, fieldName)
			sourceItems, ok := source.([]interface{})
			if ok {
				destinationItems := make([]MappedData, len(sourceItems))
				for i := range sourceItems {
					destinationItems[i] = make(MappedData)
				}

				md[fieldName] = destinationItems

				destinationPointers := make([]*MappedData, len(destinationItems))
				for i := range destinationItems {
					destinationPointers[i] = &destinationItems[i]
				}

				copyPairs = append(copyPairs, newCopyPair(sourceItems, destinationPointers))
				continue
			}

			destinationItem := make(MappedData)
			md[fieldName] = destinationItem
			copyPairs = append(copyPairs, newSingleCopyPair(source, &destinationItem))
		}
	}

	newConfigureMap := newConfigureMap(cm.resource, copyPairs...)
	return &newConfigureMap
}

func mapChildSlice(r *Resource, md *MappedData, fieldName string, sourceItems []interface{}) *ConfigureMap {
	destinationItems := make([]MappedData, len(sourceItems))
	for i := range sourceItems {
		destinationItems[i] = make(MappedData)
	}

	(*md)[fieldName] = destinationItems

	destinationPointers := make([]*MappedData, len(destinationItems))
	for i := range destinationItems {
		destinationPointers[i] = &destinationItems[i]
	}

	copyPairs := []copyPair{{sourceItems, destinationPointers}}

	csm := newConfigureMap(r, copyPairs...)
	return &csm
}

func mapChildStruct(r *Resource, md *MappedData, fieldName string, source interface{}) *ConfigureMap {
	childResourceData := make(MappedData)
	(*md)[fieldName] = childResourceData

	cm := newConfigureMap(r, newSingleCopyPair(source, &childResourceData))
	return &cm
}

func (cm *ConfigureMap) MapAll() *ConfigureMap {
	if len(cm.copyPairs[0].sourceItems) == 0 {
		return cm
	}

	firstItem := cm.copyPairs[0].sourceItems[0]

	t := reflect.TypeOf(firstItem)

	for i := 0; i < t.NumField(); i++ {
		var fieldName = t.Field(i).Name
		if !t.Field(i).IsExported() {
			continue
		}

		if slices.Contains(cm.excludedFields, fieldName) {
			continue
		}

		cm.Map(fieldName)
	}
	return cm
}

func (cm *ConfigureMap) Exclude(fieldName string) *ConfigureMap {
	cm.excludedFields = append(cm.excludedFields, fieldName)

	for _, copyPair := range cm.copyPairs {
		for i, v := range copyPair.sourceItems {
			resourceData := *copyPair.destinationItems[i]

			value := getValueByName(v, fieldName)
			value = createResourceData(value)

			delete(resourceData, fieldName)
		}
	}

	return cm
}

func (cm *ConfigureMap) EndMap() *Resource {
	return cm.resource
}
