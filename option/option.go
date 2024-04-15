package option

type Option struct {
	option string
	value  string
}

func findOption(options []Option, option string) (string, bool) {
	for _, o := range options {
		if o.option == option {
			return o.value, true
		}
	}
	return "", false
}
