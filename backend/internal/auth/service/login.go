package auth_service

import (
	"context"
	"errors"
	"fmt"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *AuthServiceImpl) LoginByUsername(ctx context.Context, input LoginByUsernameInput) (core_auth.Token, error) {
	user, err := s.repo.GetUserByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return core_auth.Token{}, fmt.Errorf("invalid credentials: %w", core_errors.ErrUnauthorized)
		}
		return core_auth.Token{}, fmt.Errorf("failed to get user: %w", err)
	}

	if err := core_auth.ComparePassword(user.Password, input.Password); err != nil {
		return core_auth.Token{}, fmt.Errorf("invalid credentials: %w", core_errors.ErrUnauthorized)
	}

	token, err := s.jwtProcessor.GenerateToken(user.ID)
	if err != nil {
		return core_auth.Token{}, fmt.Errorf("failed to generate token: %w", err)
	}
	return core_auth.Token{AccessToken: token}, nil
}

func (s *AuthServiceImpl) LoginByEmail(ctx context.Context, input LoginByEmailInput) (core_auth.Token, error) {
	user, err := s.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return core_auth.Token{}, fmt.Errorf("invalid credentials: %w", core_errors.ErrUnauthorized)
		}
		return core_auth.Token{}, fmt.Errorf("failed to get user: %w", err)
	}

	if err := core_auth.ComparePassword(user.Password, input.Password); err != nil {
		return core_auth.Token{}, fmt.Errorf("invalid credentials: %w", core_errors.ErrUnauthorized)
	}

	token, err := s.jwtProcessor.GenerateToken(user.ID)
	if err != nil {
		return core_auth.Token{}, fmt.Errorf("failed to generate token: %w", err)
	}
	return core_auth.Token{AccessToken: token}, nil
}
