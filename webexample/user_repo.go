package main

import "strings"

type userRepo struct {
	users []user
}

func (r userRepo) Search(username string) []user {
	users := make([]user, 0)
	for _, u := range r.users {
		if strings.Contains(u.Username, username) {
			users = append(users, u)
		}
	}

	return users
}

func newUserRepo() userRepo {
	users := []user{
		newUser("ajones", "ajones@aol.com"),
		newUser("sanderson", "sanderson@gmail.com"),
		newUser("mwilliams", "mwilliams@gmail.com"),
		newUser("jsmith", "jsmith@outlook.com"),
	}
	return userRepo{users}
}
