package encoding

import (
	"encoding/json"
	"encoding/xml"
	"github.com/slyjeff/rest-resource/resource"
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

		linksText := "\"_links\":" + string(links)
		if text != "{}" {
			linksText = "," + linksText
		}

		text = text[:len(text)-1] + linksText + "}"
	}

	return []byte(text), nil
}

func MarshalXml(r resource.Resource) ([]byte, error) {
	return xml.Marshal(r)
}
