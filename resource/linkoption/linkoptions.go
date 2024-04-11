package linkoption

type LinkOption struct {
	option string
	value  string
}

func Verb(name string) LinkOption {
	return LinkOption{"verb", name}
}

func Templated() LinkOption {
	return LinkOption{"isTemplated", "true"}
}

func FindVerbOption(options []LinkOption) (string, bool) {
	return findOption(options, "verb")
}

func FindTemplatedOption(options []LinkOption) bool {
	_, isTemplated := findOption(options, "isTemplated")
	return isTemplated
}

func findOption(options []LinkOption, option string) (string, bool) {
	for _, v := range options {
		if v.option == option {
			return v.value, true
		}
	}
	return "", false
}
