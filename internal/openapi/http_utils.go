package openapi

import (
	"github.com/slyjeff/rest-resource"
	"strings"
)

func MarshalDoc(headers map[string][]string, info Info, server string, resources ...resource.Resource) ([]byte, string) {
	acceptFormats, _ := headers["Accept"]

	if formatAccepted(acceptFormats, "application/json") {
		if v, err := MarshalJson(info, server, resources...); err == nil {
			return v, "application/json"
		}
		return nil, "application/json"
	}

	if v, err := MarshalYaml(info, server, resources...); err == nil {
		return v, "text/yaml"
	}

	return nil, "text/yaml"
}

func formatAccepted(acceptFormats []string, format string) bool {
	for _, af := range acceptFormats {
		if strings.Contains(af, format) {
			return true
		}
	}

	return false
}
