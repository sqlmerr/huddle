package auth_postgres_repository

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}
