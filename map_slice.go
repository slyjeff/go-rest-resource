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
		name = fieldName
	} else {
		name = mapOptions.Name
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

			rd.AddData(fieldName, destinationItems)

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
		var fieldName = t.Field(i).Name
		if slices.Contains(csm.excludedFields, fieldName) {
			continue
		}

		csm.MapWithOptions(fieldName, MapOptions{})
	}

	return csm
}

func (csm *ConfigureSliceMap) Exclude(fieldName string) ChildMapper {
	csm.excludedFields = append(csm.excludedFields, fieldName)

	for _, copyPair := range csm.copyPairs {
		for i, v := range copyPair.sourceItems {
			resourceData := copyPair.destinationItems[i]

			value := getValueByName(v, fieldName)
			value = createResourceData(value)

			delete(resourceData.Values, fieldName)
		}
	}

	return csm
}

func (csm *ConfigureSliceMap) EndMap() *Resource {
	return csm.resource
}
