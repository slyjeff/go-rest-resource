package main

type user struct {
	Username string `form:"username"`
	Email    string `form:"email"`
}

func newUser(userName, email string) user {
	return user{userName, email}
}
