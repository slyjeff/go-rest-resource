package json

import (
	"encoding/json"
	"github.com/slyjeff/rest-resource/resource"
)

func Marshal(r resource.Resource) ([]byte, bool) {
	if j, err := json.Marshal(r.Values); err == nil {
		return j, true
	}
	return make([]byte, 0), false
}
