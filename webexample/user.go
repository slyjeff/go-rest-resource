package main

type user struct {
	username string
	email    string
}

func newUser(userName, email string) user {
	return user{userName, email}
}
