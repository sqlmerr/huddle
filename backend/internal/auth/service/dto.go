package auth_service

import (
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

type LoginByUsernameInput struct {
	Username string
	Password string
}

type LoginByEmailInput struct {
	Email    string
	Password string
}

type ChangePasswordInput struct {
	UserID      uuid.UUID
	OldPassword string
	NewPassword string
}

func (i *ChangePasswordInput) Validate() error {
	if i.NewPassword == "" {
		return fmt.Errorf("new password must be non-null: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
