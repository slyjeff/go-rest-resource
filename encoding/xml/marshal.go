package xml

import (
	"encoding/xml"
	"github.com/slyjeff/rest-resource/resource"
)

func Marshal(r resource.Resource) ([]byte, error) {
	return xml.Marshal(r)
}
