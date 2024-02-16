package encoding

import (
	"encoding/xml"
	"github.com/slyjeff/rest-resource/resource"
	"strings"
)

func MarshalResource(r resource.Resource, headers map[string][]string) (string, string) {
	acceptFormats, _ := headers["Accept"]

	if formatAccepted(acceptFormats, "text/html") {
		if v, err := MarshalHtml(r); err == nil {
			return string(v), "text/html"
		}
		return "", "text/html"
	}

	if formatAccepted(acceptFormats, "application/xml") {
		if v, err := MarshalXml(r); err == nil {
			return xml.Header + string(v), "application/xml"
		}
		return "", "application/xml"
	}

	if v, err := MarshalJson(r); err == nil {
		return string(v), "application/json"
	}
	return "", "application/json"
}

func formatAccepted(acceptFormats []string, format string) bool {
	for _, af := range acceptFormats {
		if strings.Contains(af, format) {
			return true
		}
	}

	return false
}
