package auth_service

import (
	"context"
	"fmt"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *AuthService) Register(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}

	existsByUsername, err := s.repo.UserExistsByUsername(ctx, user.Username)
	if err != nil {
		return domain.User{}, fmt.Errorf("check user existence by username: %w", err)
	}
	if existsByUsername == true {
		return domain.User{}, fmt.Errorf("username '%s' already occupied: %w", user.Username, core_errors.ErrConflict)
	}

	existsByEmail, err := s.repo.UserExistsByEmail(ctx, user.Email)
	if err != nil {
		return domain.User{}, fmt.Errorf("check user existence by email: %w", err)
	}
	if existsByEmail == true {
		return domain.User{}, fmt.Errorf("email '%s' already occupied: %w", user.Email, core_errors.ErrConflict)
	}

	hashedPassword, err := core_auth.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}
	user.Password = hashedPassword

	user, err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("register: %w", err)
	}

	return user, nil
}
