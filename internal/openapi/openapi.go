package openapi

type openApi struct {
	Openapi string          `json:"openapi"`
	Info    Info            `json:"info"`
	Paths   map[string]Path `json:"paths"`
}

type Info struct {
	Title          string            `json:"title"`
	Description    string            `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string            `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        map[string]string `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License          `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string            `json:"version"`
}

type License struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Url  string `json:"url,omitempty" yaml:"url,omitempty"`
}

type Path struct {
}
