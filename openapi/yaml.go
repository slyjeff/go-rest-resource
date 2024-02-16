package openapi

import (
	"github.com/slyjeff/rest-resource/resource"
	"gopkg.in/yaml.v3"
)

func MarshalYaml(info Info, r ...resource.Resource) ([]byte, error) {
	doc := openApi{
		"3.0.3",
		info,
		make(map[string]Path),
	}

	b, err := yaml.Marshal(doc)
	return b, err
}
