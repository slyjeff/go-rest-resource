package option

func Verb(name string) Option {
	return Option{"verb", name}
}

func Templated() Option {
	return Option{"isTemplated", "true"}
}

func FindVerbOption(options []Option) (string, bool) {
	return findOption(options, "verb")
}

func FindTemplatedOption(options []Option) bool {
	_, isTemplated := findOption(options, "isTemplated")
	return isTemplated
}
