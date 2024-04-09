package linkoption

type LinkOption struct {
	option string
	value  string
}

func Verb(name string) LinkOption {
	return LinkOption{"verb", name}
}

func FindVerbOption(options []LinkOption) (string, bool) {
	return findOption(options, "verb")
}

func findOption(options []LinkOption, option string) (string, bool) {
	for _, v := range options {
		if v.option == option {
			return v.value, true
		}
	}
	return "", false
}
