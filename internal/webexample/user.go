package main

type user struct {
	Id       int
	Username string `form:"username"`
	Email    string `form:"email"`
	IsActive bool   `form:"is_active"`
}
