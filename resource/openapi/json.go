package openapi

import (
	"encoding/json"
	"github.com/slyjeff/rest-resource/resource"
)

func MarshalJson(info Info, r ...resource.Resource) ([]byte, error) {
	doc := openApi{
		"3.0.3",
		info,
		make(map[string]Path),
	}

	return json.MarshalIndent(doc, "", "  ")
}
