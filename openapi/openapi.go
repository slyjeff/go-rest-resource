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

		if _, ok := doc.Components.Schemas[r.Schema]; !ok {
			doc.Components.Schemas[r.Schema] = newSchemaFromResource(r)
		}

		for _, link := range r.Links {
			if link.Verb == "GET" || len(link.Parameters) == 0 || link.ResponseSchema == "" {
				continue
			}

			bodySchema := link.Verb + link.ResponseSchema
			if _, ok := doc.Components.Schemas[bodySchema]; !ok {
				doc.Components.Schemas[bodySchema] = newSchemaFromParameters(link.Parameters)
			}
		}
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
		parameters := getPathParameters(link.Href)
		if len(parameters) > 0 {
			path["parameters"] = parameters
		}
		openApi.Paths[link.Href] = path
	}

	verb := strings.ToLower(link.Verb)

	queryParameters := getQueryParameters(link)
	operation, ok := path[verb]
	if !ok {
		bodySchema := ""
		if link.Verb != "GET" && link.ResponseSchema != "" && len(link.Parameters) > 0 {
			bodySchema = link.Verb + link.ResponseSchema
		}

		path[verb] = newOperation(link.ResponseCodes, formatSummary(summary), link.ResponseSchema, queryParameters, bodySchema)
	} else if len(queryParameters) > 0 {
		if o, ok := operation.(Operation); ok {
			o.QueryParameters = queryParameters
		}
	}
}

func getPathParameters(url string) []Parameter {
	parameters := make([]Parameter, 0)

	re := regexp.MustCompile("{[a-zA-Z0-9]*}")
	parameterNames := re.FindAllString(url, -1)
	if len(parameterNames) == 0 {
		return parameters
	}

	for _, parameter := range parameterNames {
		parameter = parameter[1 : len(parameter)-1]
		parameters = append(parameters, Parameter{parameter, "path", true, newIntSchema()})
	}

	return parameters
}

func getQueryParameters(link resource.Link) []Parameter {
	parameters := make([]Parameter, 0)

	if link.Verb != "GET" {
		return parameters
	}

	for _, parameter := range link.Parameters {
		parameters = append(parameters, Parameter{parameter.Name, "query", false, newSchemaFromDataType(parameter.DataType)})
	}

	return parameters
}

func newSchemaFromDataType(dataType string) Schema {
	switch strings.ToLower(dataType) {
	case "int32":
		return newInt32Schema()
	case "int", "int64", "number":
		return newIntSchema()
	case "float", "float32", "float64":
		return newFloatSchema()
	case "bool", "boolean":
		return newBoolSchema()
	default:
		return newStringSchema()
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
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	In       string `json:"in,omitempty" yaml:"in,omitempty"`
	Required bool   `json:"required" yaml:"required"`
	Schema   Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

type Operation struct {
	Description     string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Responses       map[string]DataObject `json:"responses,omitempty" yaml:"responses,omitempty"`
	QueryParameters []Parameter           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody     *DataObject           `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
}

func newOperation(codes []int, description string, schema string, queryParameters []Parameter, bodySchema string) Operation {
	responses := make(map[string]DataObject)
	for _, code := range codes {
		content := make(map[string]Content)
		if code >= 200 && code <= 202 && schema != "" {
			responseContent := newResponseContent(schema)
			content["application/json"] = responseContent
			content["application/xml"] = responseContent
		}
		responses[fmt.Sprintf("%v", code)] = DataObject{http.StatusText(code), content}
	}

	var requestBody *DataObject = nil
	if bodySchema != "" {
		requestBody = &DataObject{Content: make(map[string]Content)}
		requestBodyContent := newResponseContent(bodySchema)
		requestBody.Content["application/json"] = requestBodyContent
		requestBody.Content["application/xml"] = requestBodyContent
		requestBody.Content["application/x-www-form-urlencoded"] = requestBodyContent
	}

	return Operation{description, responses, queryParameters, requestBody}
}

type DataObject struct {
	Description string             `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]Content `json:"content,omitempty" yaml:"content,omitempty"`
}

type Content struct {
	Schema RefSchema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

func newResponseContent(schema string) Content {
	return Content{RefSchema{Ref: "#/components/schemas/" + schema}}
}

type RefSchema struct {
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

func newSchemaFromParameters(parameters []resource.LinkParameter) Schema {
	schema := Schema{"object", "", make(map[string]Schema), nil}

	for _, parameter := range parameters {
		schema.Properties[parameter.Name] = newSchemaFromDataType(parameter.DataType)
	}

	return schema
}

func newSchemaFromValue(value interface{}) Schema {
	if fd, ok := value.(resource.FormattedData); ok {
		value = fd.Value
	}

	switch value.(type) {
	case int:
		return newIntSchema()
	case float64:
		return newFloatSchema()
	case bool:
		return newBoolSchema()
	default:
		return newStringSchema()
	}
}

func newInt32Schema() Schema {
	return Schema{"integer", "int32", make(map[string]Schema), nil}
}

func newIntSchema() Schema {
	return Schema{"integer", "int64", make(map[string]Schema), nil}
}

func newFloatSchema() Schema {
	return Schema{"number", "float", make(map[string]Schema), nil}
}

func newBoolSchema() Schema {
	return Schema{"boolean", "", make(map[string]Schema), nil}
}

func newStringSchema() Schema {
	return Schema{"string", "", make(map[string]Schema), nil}
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
