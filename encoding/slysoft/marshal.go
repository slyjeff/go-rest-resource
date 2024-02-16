package slysoft

import (
	"encoding/json"
	"encoding/xml"
	"github.com/slyjeff/rest-resource/resource"
)

func MarshalJson(r resource.Resource) ([]byte, error) {
	return json.Marshal(r.Values)
}

func MarshalXml(r resource.Resource) ([]byte, error) {
	return xml.Marshal(r)
}
