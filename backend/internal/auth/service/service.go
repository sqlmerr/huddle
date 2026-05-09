package auth_service

import (
	"context"

	"github.com/google/uuid"
	auth_postgres_repository "github.com/sqlmerr/huddle/backend/internal/auth/repository/postgres"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type AuthServiceImpl struct {
	repo         auth_postgres_repository.AuthRepository
	jwtProcessor JWTProcessor
}

type JWTProcessor interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
}

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (domain.User, error)
	LoginByUsername(ctx context.Context, input LoginByUsernameInput) (core_auth.Token, error)
	LoginByEmail(ctx context.Context, input LoginByEmailInput) (core_auth.Token, error)
	ChangePassword(ctx context.Context, input ChangePasswordInput) error
}

func NewAuthService(repo auth_postgres_repository.AuthRepository, jwtProcessor JWTProcessor) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo, jwtProcessor: jwtProcessor}
}
