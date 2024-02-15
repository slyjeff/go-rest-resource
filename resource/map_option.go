package resource

type MapOptions struct {
	Name           string
	FormatCallback FormatDataCallback
}

type MapOption struct {
	option string
	value  string
}

func RenameField(name string) MapOption {
	return MapOption{"name", name}
}

func FormatField(formatString string) MapOption {
	return MapOption{"format", formatString}
}

func FindNameOption(options []MapOption) (string, bool) {
	return FindOption(options, "name")
}

func FindFormatOption(options []MapOption) (string, bool) {
	return FindOption(options, "format")
}

func FindOption(options []MapOption, option string) (string, bool) {
	for _, v := range options {
		if v.option == option {
			return v.value, true
		}
	}
	return "", false
}
