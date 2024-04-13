package main

type user struct {
	Username string `form:"username" query:"username"`
	Email    string `form:"email" query:"email"`
}

func newUser(userName, email string) user {
	return user{userName, email}
}
