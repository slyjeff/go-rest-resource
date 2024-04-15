package main

import "strings"

type userRepo struct {
	users []user
}

func newUserRepo() userRepo {
	userRepo := userRepo{[]user{
		{1, "ajones", "ajones@aol.com"},
		{2, "sanderson", "sanderson@gmail.com"},
		{3, "mwilliams", "mwilliams@gmail.com"},
		{4, "jsmith", "jsmith@outlook.com"},
	}}

	return userRepo
}

func (r *userRepo) Search(username string) []user {
	users := make([]user, 0)
	for _, u := range r.users {
		if strings.Contains(u.Username, username) {
			users = append(users, u)
		}
	}

	return users
}

func (r *userRepo) Add(u *user) {
	u.id = r.GetUniqueId()
	r.users = append(r.users, *u)
}

func (r *userRepo) GetUniqueId() int {
	id := 1
	for _, u := range r.users {
		if u.id >= id {
			id = u.id + 1
		}
	}

	return id
}

func (r *userRepo) GetById(id int) (*user, bool) {
	for _, u := range r.users {
		if u.id == id {
			return &u, true
		}
	}

	return nil, false
}

func (r *userRepo) Delete(id int) bool {
	foundIndex := -1
	for i, u := range r.users {
		if u.id == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return false
	}

	r.users = append(r.users[:foundIndex], r.users[foundIndex+1:]...)
	return true
}
