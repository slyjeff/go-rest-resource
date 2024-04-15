package encoding

import (
	"encoding/json"
	"encoding/xml"
	"github.com/slyjeff/rest-resource"
	"strings"
)

func MarshalJson(r resource.Resource) ([]byte, error) {
	values, err := json.Marshal(r.Values)
	if err != nil {
		return nil, err
	}

	text := string(values)
	if text == "null" {
		text = "{}"
	}

	if len(r.Links) > 0 {
		var links []byte
		links, err = json.Marshal(r.Links)
		if err != nil {
			return nil, err
		}

		text = addToJson(text, "_links", string(links))
	}

	if len(r.Embedded) > 0 {
		embeddedJson := "{}"
		for name, embedded := range r.Embedded {
			if embeddedResource, ok := embedded.(resource.Resource); ok {
				if resourceJson, err := MarshalJson(embeddedResource); err == nil {
					embeddedJson = addToJson(embeddedJson, name, string(resourceJson))
				}
			} else if embeddedResourceList, ok := embedded.([]resource.Resource); ok {
				resourceJsonList := make([]string, 0)
				for _, embeddedResource := range embeddedResourceList {
					if resourceJson, err := MarshalJson(embeddedResource); err == nil {
						resourceJsonList = append(resourceJsonList, string(resourceJson))
					}
				}
				var jsonArray = "[" + strings.Join(resourceJsonList, ",") + "]"
				embeddedJson = addToJson(embeddedJson, name, jsonArray)
			}
		}

		text = addToJson(text, "_embedded", embeddedJson)
	}

	return []byte(text), nil
}

func addToJson(json, name, value string) string {
	nameValue := "\"" + name + "\":" + value
	if json == "{}" {
		return json[:len(json)-1] + nameValue + "}"
	}
	return json[:len(json)-1] + "," + nameValue + "}"
}

func MarshalXml(r resource.Resource) ([]byte, error) {
	return xml.Marshal(r)
}
