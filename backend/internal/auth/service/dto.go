package auth_service

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

func (u *RegisterInput) Validate() error {
	usernameLen := len([]rune(u.Username))
	if usernameLen < 3 || usernameLen > 32 {
		return fmt.Errorf("invalid `Username` length: %d: %w", usernameLen, core_errors.ErrInvalidArgument)
	}

	if !usernameRegex.MatchString(u.Username) {
		return fmt.Errorf("invalid `Username` format: must contain only latin letters, digits, and underscores: %w", core_errors.ErrInvalidArgument)
	}

	if len(u.Email) == 0 {
		return fmt.Errorf("invalid `Email`: cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if !emailRegex.MatchString(u.Email) {
		return fmt.Errorf("invalid `Email` format: %w", core_errors.ErrInvalidArgument)
	}

	passwordLen := len([]rune(u.Password))
	if passwordLen < 1 {
		return fmt.Errorf("invalid `Password` length: %d: must be at least 1 character: %w", passwordLen, core_errors.ErrInvalidArgument)
	}

	return nil
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
