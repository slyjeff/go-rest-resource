package encoding

import (
	"bytes"
	"encoding/json"
	"github.com/slyjeff/rest-resource/resource"
	"html/template"
	"reflect"
	"strings"
)

func MarshalHtml(r resource.Resource) ([]byte, error) {
	t := template.New("ResourceTemplate")

	t = t.Funcs(template.FuncMap{
		"FormatValue": func(i interface{}) interface{} {
			fd, ok := i.(resource.FormattedData)
			if ok {
				return fd.FormattedString()
			}

			isValue := true
			kind := reflect.TypeOf(i).Kind()
			if kind == reflect.Slice || kind == reflect.Array {
				isValue = false
			} else {
				_, ok := i.(resource.MappedData)
				isValue = !ok
			}

			if isValue {
				return i
			}

			if j, err := json.Marshal(i); err == nil {
				return string(j)
			}

			return ""
		},
		"SeparateListOfValues": func(s string) []string {
			return strings.Split(s, ",")
		},
	})

	var err error
	t, err = t.Parse(resourceHtml)
	if err != nil {
		return make([]byte, 0), err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, r); err != nil {
		return make([]byte, 0), err
	}
	return buf.Bytes(), nil
}
