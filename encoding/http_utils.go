package encoding

import (
	"encoding/xml"
	"github.com/slyjeff/rest-resource/encoding/slysoft"
	"github.com/slyjeff/rest-resource/resource"
)

func ForAcceptHeader(r resource.Resource, headers map[string][]string) (string, string) {
	if formatAccepted(headers, "application/xml") {
		if v, err := slysoft.MarshalXml(r); err == nil {
			return xml.Header + string(v), "application/xml"
		}
		return "", "application/xml"
	}

	if v, err := slysoft.MarshalJson(r); err == nil {
		return string(v), "application/json"
	}
	return "", "application/json"
}

func formatAccepted(headers map[string][]string, format string) bool {
	values, ok := headers["Accept"]
	if !ok {
		return false
	}

	for _, v := range values {
		if v == format {
			return true
		}
	}

	return false
}
