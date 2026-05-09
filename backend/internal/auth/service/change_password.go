package auth_service

import (
	"context"
	"fmt"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *AuthServiceImpl) ChangePassword(ctx context.Context, input ChangePasswordInput) error {
	if err := input.Validate(); err != nil {
		return fmt.Errorf("validate input: %w", err)
	}

	user, err := s.repo.GetUser(ctx, input.UserID)
	if err != nil {
		return fmt.Errorf("get user by id='%s': %w", input.UserID.String(), err)
	}

	if err := core_auth.ComparePassword(user.Password, input.OldPassword); err != nil {
		return fmt.Errorf("password not match: %w", core_errors.ErrUnprocessableEntity)
	}

	hashedPassword, err := core_auth.HashPassword(input.NewPassword)
	if err != nil {
		return fmt.Errorf("hash password: %w", core_errors.ErrInternalServerError)
	}

	user.Password = hashedPassword
	err = s.repo.SaveUser(ctx, user)
	if err != nil {
		return fmt.Errorf("save user: %w: %w", core_errors.ErrInternalServerError, err)
	}

	return nil
}
