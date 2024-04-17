package main

import (
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type userRepo struct {
	users []user
}

func newUserRepo() userRepo {
	userRepo := userRepo{[]user{
		{uuid.New(), "ajones", "ajones@aol.com", true},
		{uuid.New(), "sanderson", "sanderson@gmail.com", false},
		{uuid.New(), "mwilliams", "mwilliams@gmail.com", true},
		{uuid.New(), "jsmith", "jsmith@outlook.com", true},
	}}

	return userRepo
}

func (r *userRepo) Search(userSearch userSearch) []user {
	users := make([]user, 0)
	for _, u := range r.users {
		if canAdd(u, userSearch) {
			users = append(users, u)
		}
	}

	return users
}

func canAdd(user user, userSearch userSearch) bool {
	if userSearch.Username != "" && !strings.Contains(user.Username, userSearch.Username) {
		return false
	}

	if userSearch.IsActive != "" {
		isActive, err := strconv.ParseBool(userSearch.IsActive)
		if err != nil {
			return false
		}
		return user.IsActive == isActive
	}

	return true
}

func (r *userRepo) Add(u *user) {
	u.Id = uuid.New()
	r.users = append(r.users, *u)
}

func (r *userRepo) GetById(id uuid.UUID) (*user, bool) {
	for _, u := range r.users {
		if u.Id == id {
			return &u, true
		}
	}

	return nil, false
}

func (r *userRepo) Delete(id uuid.UUID) bool {
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
