package main

type user struct {
	id       int
	Username string `form:"username"`
	Email    string `form:"email"`
}
