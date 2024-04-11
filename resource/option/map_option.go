package option

func Rename(name string) Option {
	return Option{"name", name}
}

func Format(formatString string) Option {
	return Option{"format", formatString}
}

func FindNameOption(options []Option) (string, bool) {
	return findOption(options, "name")
}

func FindFormatOption(options []Option) (string, bool) {
	return findOption(options, "format")
}
