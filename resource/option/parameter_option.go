package option

func Default(name string) Option {
	return Option{"default", name}
}

func FindDefaultOption(options []Option) (string, bool) {
	return findOption(options, "default")
}
