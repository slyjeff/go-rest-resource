package main

import "strings"

type userRepo struct {
	users []user
}

func newUserRepo() userRepo {
	userRepo := userRepo{[]user{
		{1, "ajones", "ajones@aol.com", true},
		{2, "sanderson", "sanderson@gmail.com", false},
		{3, "mwilliams", "mwilliams@gmail.com", true},
		{4, "jsmith", "jsmith@outlook.com", true},
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
	u.Id = r.GetUniqueId()
	r.users = append(r.users, *u)
}

func (r *userRepo) GetUniqueId() int {
	id := 1
	for _, u := range r.users {
		if u.Id >= id {
			id = u.Id + 1
		}
	}

	return id
}

func (r *userRepo) GetById(id int) (*user, bool) {
	for _, u := range r.users {
		if u.Id == id {
			return &u, true
		}
	}

	return nil, false
}

func (r *userRepo) Delete(id int) bool {
	foundIndex := -1
	for i, u := range r.users {
		if u.Id == id {
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
