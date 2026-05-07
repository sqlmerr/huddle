package auth_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type AuthService struct {
	repo         AuthRepository
	jwtProcessor JWTProcessor
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UserExistsByUsername(ctx context.Context, username string) (bool, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
}

type JWTProcessor interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
}

func NewAuthService(repo AuthRepository, jwtProcessor JWTProcessor) *AuthService {
	return &AuthService{repo: repo, jwtProcessor: jwtProcessor}
}
