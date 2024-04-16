package openapi

import (
	"fmt"
	resource "github.com/slyjeff/rest-resource"
	"net/http"
	"regexp"
	"strings"
)

type Info struct {
	Title          string            `json:"title"`
	Description    string            `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string            `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        map[string]string `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License          `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string            `json:"version"`
}

func newOpenApi(info Info, server string, resources []resource.Resource) openApi {
	doc := openApi{
		"3.0.3",
		info,
		[]Server{{server}},
		make(map[string]Path),
		Components{make(map[string]Schema)},
	}

	for _, r := range resources {
		for linkName, link := range r.Links {
			if linkName == "self" {
				linkName = "Get" + r.Schema
			}
			doc.addPath(*link, linkName)
		}

		if r.Schema == "" {
			continue
		}

		if _, ok := doc.Components.Schemas[r.Schema]; ok {
			continue
		}

		doc.Components.Schemas[r.Schema] = newSchemaFromResource(r)
	}

	return doc
}

type openApi struct {
	Openapi    string          `json:"openapi" yaml:"openapi"`
	Info       Info            `json:"info" yaml:"info"`
	Servers    []Server        `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths      map[string]Path `json:"paths" yaml:"paths"`
	Components Components      `json:"components,omitempty" yaml:"components,omitempty"`
}

func (openApi *openApi) addPath(link resource.Link, summary string) {
	path, ok := openApi.Paths[link.Href]
	if !ok {
		path = make(Path)
		re := regexp.MustCompile("{[a-zA-Z0-9]*}")
		parameterNames := re.FindAllString(link.Href, -1)
		if len(parameterNames) > 0 {
			parameters := make([]Parameter, 0)
			for _, parameter := range parameterNames {
				parameter = parameter[1 : len(parameter)-1]
				parameters = append(parameters, Parameter{parameter, "path", true, newIntParameterSchema()})
			}
			path["parameters"] = parameters
		}
		openApi.Paths[link.Href] = path
	}

	verb := strings.ToLower(link.Verb)
	if _, ok := path[verb]; !ok {
		path[verb] = newOperation(link.ResponseCodes, formatSummary(summary), link.Schema)
	}
}

func formatSummary(s string) string {
	s = separateWords(s)
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

func separateWords(s string) string {
	re := regexp.MustCompile(`([A-Z]+)`)
	s = re.ReplaceAllString(s, ` $1`)
	s = strings.Trim(s, " ")
	return s
}

type Server struct {
	Url string `json:"url" yaml:"url"`
}

type License struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Url  string `json:"url,omitempty" yaml:"url,omitempty"`
}

type Path map[string]interface{}

type Parameter struct {
	Name     string          `json:"name,omitempty" yaml:"name,omitempty"`
	In       string          `json:"in,omitempty" yaml:"in,omitempty"`
	Required bool            `json:"required" yaml:"required"`
	Schema   ParameterSchema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

type ParameterSchema struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}

func newIntParameterSchema() ParameterSchema {
	return ParameterSchema{"integer", "int64"}
}

type Operation struct {
	Description string              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty" yaml:"responses,omitempty"`
}

func newOperation(codes []int, description string, schema string) Operation {
	responses := make(map[string]Response)
	for _, code := range codes {
		content := make(map[string]ResponseContent)
		if code >= 200 && code <= 202 && schema != "" {
			responseContent := newResponseContent(schema)
			content["application/json"] = responseContent
			content["application/xml"] = responseContent
		}
		responses[fmt.Sprintf("%v", code)] = Response{http.StatusText(code), content}
	}

	return Operation{description, responses}
}

type Response struct {
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]ResponseContent `json:"content,omitempty" yaml:"content,omitempty"`
}

type ResponseContent struct {
	Schema ResponseSchema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

func newResponseContent(schema string) ResponseContent {
	return ResponseContent{ResponseSchema{Ref: "#/components/schemas/" + schema}}
}

type ResponseSchema struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Components struct {
	Schemas map[string]Schema `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

type Schema struct {
	Type       string            `json:"type,omitempty" yaml:"type,omitempty"`
	Format     string            `json:"format,omitempty" yaml:"format,omitempty"`
	Properties map[string]Schema `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items      *Schema           `json:"items,omitempty" yaml:"items,omitempty"`
}

func newSchemaFromResource(r resource.Resource) Schema {
	schema := Schema{"object", "", make(map[string]Schema), nil}

	for name, value := range r.Values {
		schema.Properties[name] = newSchemaFromValue(value)
	}

	if len(r.Embedded) > 0 {
		schema.Properties["_embedded"] = newSchemaFromEmbedded(r.Embedded)
	}

	return schema
}

func newSchemaFromValue(value interface{}) Schema {
	if fd, ok := value.(resource.FormattedData); ok {
		value = fd.Value
	}

	switch value.(type) {
	case int:
		return Schema{"integer", "int65", make(map[string]Schema), nil}
	case float64:
		return Schema{"number", "float", make(map[string]Schema), nil}
	case bool:
		return Schema{"boolean", "", make(map[string]Schema), nil}
	default:
		return Schema{"string", "", make(map[string]Schema), nil}
	}
}

func newSchemaFromEmbedded(embeddedResources resource.EmbeddedResources) Schema {
	schema := Schema{"object", "", make(map[string]Schema), nil}

	for name, embedded := range embeddedResources {
		if embeddedResource, ok := embedded.(resource.Resource); ok {
			schema.Properties[name] = newSchemaFromResource(embeddedResource)
		} else if embeddedList, ok := embedded.([]resource.Resource); ok {
			schema.Properties[name] = newSchemaFromEmbeddedList(embeddedList)
		}
	}

	return schema
}

func newSchemaFromEmbeddedList(resources []resource.Resource) Schema {
	schema := Schema{"array", "", make(map[string]Schema), nil}
	if len(resources) == 0 {
		return schema
	}

	items := newSchemaFromResource(resources[0])
	schema.Items = &items
	return schema
}
