package openapi

import (
	resource "github.com/slyjeff/rest-resource"
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
	}

	for _, r := range resources {
		for linkName, link := range r.Links {
			if linkName == "self" {
				linkName = "Get" + r.Name
			}
			doc.addPath(*link, linkName)
		}
	}

	return doc
}

type openApi struct {
	Openapi string          `json:"openapi"`
	Info    Info            `json:"info"`
	Servers []Server        `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths   map[string]Path `json:"paths"`
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
		path[verb] = newOperation(formatSummary(summary))
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

func newOperation(description string) Operation {
	responses := make(map[string]Response)
	responses["200"] = Response{"successful operation"}
	responses["500"] = Response{"internal server error"}
	return Operation{description, responses}
}

type Response struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
