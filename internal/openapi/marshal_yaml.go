package openapi

import (
	"github.com/slyjeff/rest-resource"
	"gopkg.in/yaml.v3"
)

func MarshalYaml(info Info, server string, r ...resource.Resource) ([]byte, error) {
	doc := newOpenApi(info, server, r)

	return yaml.Marshal(doc)
}
