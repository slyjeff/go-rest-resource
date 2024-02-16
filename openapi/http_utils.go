package openapi

import (
	"github.com/slyjeff/rest-resource/resource"
	"strings"
)

func MarshalDoc(headers map[string][]string, info Info, resources ...resource.Resource) (string, string) {
	acceptFormats, _ := headers["Accept"]

	if formatAccepted(acceptFormats, "application/json") {
		if v, err := MarshalJson(info, resources...); err == nil {
			return string(v), "application/json"
		}
		return "", "application/json"
	}

	if v, err := MarshalYaml(info, resources...); err == nil {
		return string(v), "text/yaml"
	}
	return "", "text/yaml"
}

func formatAccepted(acceptFormats []string, format string) bool {
	for _, af := range acceptFormats {
		if strings.Contains(af, format) {
			return true
		}
	}

	return false
}
