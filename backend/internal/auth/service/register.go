package auth_service

import (
	"context"
	"fmt"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *AuthServiceImpl) Register(ctx context.Context, input RegisterInput) (domain.User, error) {
	hashedPassword, err := core_auth.HashPassword(input.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	userDomain := domain.NewUserUninitialized(input.Username, input.Email, hashedPassword)

	if err := userDomain.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}

	existsByUsername, err := s.repo.UserExistsByUsername(ctx, input.Username)
	if err != nil {
		return domain.User{}, fmt.Errorf("check user existence by username: %w", err)
	}
	if existsByUsername == true {
		return domain.User{}, fmt.Errorf("username '%s' already occupied: %w", input.Username, core_errors.ErrConflict)
	}

	existsByEmail, err := s.repo.UserExistsByEmail(ctx, input.Email)
	if err != nil {
		return domain.User{}, fmt.Errorf("check user existence by email: %w", err)
	}
	if existsByEmail == true {
		return domain.User{}, fmt.Errorf("email '%s' already occupied: %w", input.Email, core_errors.ErrConflict)
	}

	user, err := s.repo.CreateUser(ctx, userDomain)
	if err != nil {
		return domain.User{}, fmt.Errorf("register user: %w", err)
	}

	return user, nil
}
