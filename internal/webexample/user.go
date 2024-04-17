package main

import "github.com/google/uuid"

type user struct {
	Id       uuid.UUID `json:"id"`
	Username string    `form:"username"`
	Email    string    `form:"email"`
	IsActive bool      `form:"is_active"`
}
