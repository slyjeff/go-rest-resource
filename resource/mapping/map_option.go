package mapping

type MapOption struct {
	option string
	value  string
}

func Rename(name string) MapOption {
	return MapOption{"name", name}
}

func Format(formatString string) MapOption {
	return MapOption{"format", formatString}
}

func FindNameOption(options []MapOption) (string, bool) {
	return findOption(options, "name")
}

func FindFormatOption(options []MapOption) (string, bool) {
	return findOption(options, "format")
}

func findOption(options []MapOption, option string) (string, bool) {
	for _, v := range options {
		if v.option == option {
			return v.value, true
		}
	}
	return "", false
}
