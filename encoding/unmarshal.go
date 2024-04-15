package encoding

import (
	"encoding/json"
	"github.com/slyjeff/rest-resource"
)

func UnmarshalJson(jsonToUnmarshal []byte) (resource.Resource, error) {
	var result map[string]interface{}
	r := resource.NewResource()
	if err := json.Unmarshal(jsonToUnmarshal, &result); err != nil {
		return r, err
	}

	for k, v := range result {
		if k == "_links" {
			addLinksToResource(&r, v.(map[string]interface{}))
			continue
		}
		r.Data(k, v)
	}

	return r, nil
}

func addLinksToResource(resource *resource.Resource, links map[string]interface{}) {
	for k, v := range links {
		link := v.(map[string]interface{})
		resource.Link(k, link["href"].(string))
	}
}
