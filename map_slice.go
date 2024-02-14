package GoRestResource

import (
	"reflect"
	"slices"
)

type copyPair struct {
	sourceItems      []interface{}
	destinationItems []ResourceData
}

type ConfigureSliceMap struct {
	resource       *Resource
	copyPairs      []copyPair
	excludedFields []string
}

func newConfigureSliceMap(r *Resource, copyPairs []copyPair) ConfigureSliceMap {
	return ConfigureSliceMap{r, copyPairs, make([]string, 0)}
}

func (csm *ConfigureSliceMap) Map(fieldName string) ChildMapper {
	csm.MapWithOptions(fieldName, MapOptions{})

	return csm
}

func (csm *ConfigureSliceMap) MapWithOptions(fieldName string, mapOptions MapOptions) ChildMapper {
	var name string
	if mapOptions.Name == "" {
		name = makeCamelCase(fieldName)
	} else {
		name = makeCamelCase(mapOptions.Name)
	}

	for _, copyPair := range csm.copyPairs {
		for i, v := range copyPair.sourceItems {
			resourceData := copyPair.destinationItems[i]

			value := getValueByName(v, fieldName)
			value = createResourceData(value)
			if mapOptions.FormatCallback != nil {
				value = FormattedData{value, mapOptions.FormatCallback}
			}

			resourceData.Values[name] = value
		}
	}

	return csm
}

func (csm *ConfigureSliceMap) MapChild(fieldName string) ChildMapper {
	copyPairs := make([]copyPair, 0)

	for _, cp := range csm.copyPairs {
		for i, v := range cp.sourceItems {
			value := getValueByName(v, fieldName)
			sourceItems, ok := value.([]interface{})
			if !ok {
				continue
			}

			rd := cp.destinationItems[i]
			destinationItems := make([]ResourceData, len(sourceItems))
			for i := range sourceItems {
				destinationItems[i] = ResourceData{make(map[string]interface{})}
			}

			if rd.Values == nil {
				rd.Values = make(map[string]interface{})
			}

			name := makeCamelCase(fieldName)
			rd.Values[name] = destinationItems

			copyPairs = append(copyPairs, copyPair{sourceItems, destinationItems})
		}
	}

	newConfigureSliceMap := newConfigureSliceMap(csm.resource, copyPairs)
	return &newConfigureSliceMap
}

func (csm *ConfigureSliceMap) MapAll() ChildMapper {
	if len(csm.copyPairs[0].sourceItems) == 0 {
		return csm
	}

	firstItem := csm.copyPairs[0].sourceItems[0]

	t := reflect.TypeOf(firstItem)

	for i := 0; i < t.NumField(); i++ {
		var fieldName = makeCamelCase(t.Field(i).Name)
		if slices.Contains(csm.excludedFields, fieldName) {
			continue
		}

		csm.MapWithOptions(t.Field(i).Name, MapOptions{})
	}

	return csm
}

func (csm *ConfigureSliceMap) Exclude(fieldName string) ChildMapper {
	name := makeCamelCase(fieldName)
	csm.excludedFields = append(csm.excludedFields, name)

	for _, copyPair := range csm.copyPairs {
		for i, v := range copyPair.sourceItems {
			resourceData := copyPair.destinationItems[i]

			value := getValueByName(v, fieldName)
			value = createResourceData(value)

			delete(resourceData.Values, name)
		}
	}

	return csm
}

func (csm *ConfigureSliceMap) EndMap() *Resource {
	return csm.resource
}

//
//type ConfigureSliceMapFromResource struct {
//	configureSlice
//	resource *Resource
//}
//
//func newConfigureSliceMapFromResource(r *Resource, slice []ResourceData, source []interface{}) ConfigureSliceMapFromResource {
//	return ConfigureSliceMapFromResource{newConfigureSlice(&r.ResourceData, slice, source), r}
//}
//
//type ConfigureChildSliceOfResourceMap struct {
//	configureSlice
//	parent *ConfigureResourceMap
//}
//
//func newConfigureChildSliceOfResourceMap(parent *ConfigureResourceMap, slice []ResourceData, source []interface{}) ConfigureChildSliceOfResourceMap {
//	return ConfigureChildSliceOfResourceMap{newConfigureSlice(parent.ResourceMap, slice, source), parent}
//}
//
//type ConfigureChildSliceMap struct {
//	configureSlice
//	parent *ConfigureResourceMap
//}
//
//func newConfigureChildSliceOfResourceMap(parent *ConfigureResourceMap, slice []ResourceData, source []interface{}) ConfigureChildSliceOfResourceMap {
//	return ConfigureChildSliceOfResourceMap{newConfigureSlice(parent.ResourceMap, slice, source), parent}
//}
//
//func (configuration *ConfigureSliceMapFromResource) Map(fieldName string) *ConfigureSliceMapFromResource {
//	configuration.MapWithOptions(fieldName, MapOptions{})
//
//	return configuration
//}
//
//func (configuration *ConfigureSliceMapFromResource) MapWithOptions(fieldName string, mapOptions MapOptions) *ConfigureSliceMapFromResource {
//	configuration.configureSlice.mapWithOptions(fieldName, mapOptions)
//	return configuration
//}
//
//func (configuration *ConfigureSliceMapFromResource) MapAll() *ConfigureSliceMapFromResource {
//	configuration.configureSlice.mapAll()
//	return configuration
//}
//
//func (configuration *ConfigureSliceMapFromResource) Exclude(fieldName string) *ConfigureSliceMapFromResource {
//	configuration.excludeField(fieldName)
//
//	return configuration
//}
//
//func (configuration *ConfigureSliceMapFromResource) EndMap() *Resource {
//	return configuration.resource
//}
//
//func (configuration *ConfigureChildSliceOfResourceMap) Map(fieldName string) ConfigureChildOfResourceMap {
//	configuration.MapWithOptions(fieldName, MapOptions{})
//	return configuration
//}
//
//func (configuration *ConfigureChildSliceOfResourceMap) MapWithOptions(fieldName string, mapOptions MapOptions) ConfigureChildOfResourceMap {
//	configuration.configureSlice.mapWithOptions(fieldName, mapOptions)
//	return configuration
//}
//
//func (configuration *ConfigureChildSliceOfResourceMap) MapChild(fieldName string) ConfigureChildMap {
//	slice, source := configuration.configureSlice.mapChild(fieldName)
//
//	childConfiguration = newConfigureChildSliceMap()
//	return &childConfiguration
//}
//
//func (configuration *ConfigureChildSliceOfResourceMap) MapAll() ConfigureChildOfResourceMap {
//	configuration.configureSlice.mapAll()
//	return configuration
//}
//
//func (configuration *ConfigureChildSliceOfResourceMap) EndMap() *ConfigureResourceMap {
//	return configuration.parent
//}
