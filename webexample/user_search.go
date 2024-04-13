package main

type userSearch struct {
	Username string `query:"username"`
}

func (s userSearch) Criteria() string {
	if s.Username == "" {
		return ""
	}
	return "?username=" + s.Username
}
