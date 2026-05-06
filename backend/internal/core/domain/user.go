package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(id uuid.UUID, username, email, password string, createdAt time.Time) User {
	return User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
	}
}

func NewUserUninitialized(username, email, password string) User {
	return User{
		ID:        UninitializedID,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: UninitializedTime,
	}
}
