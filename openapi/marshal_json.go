package openapi

import (
	"encoding/json"
	"github.com/slyjeff/rest-resource"
)

func MarshalJson(info Info, server string, r ...resource.Resource) ([]byte, error) {
	doc := newOpenApi(info, server, r)

	return json.MarshalIndent(doc, "", "  ")
}
