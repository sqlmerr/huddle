package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (r *UsersRepositoryImpl) GetUser(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, username, email, password, created_at
	FROM users
	WHERE id = $1
	`
	row := r.pool.QueryRow(ctx, query, userID)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Username,
		&userModel.Email,
		&userModel.Password,
		&userModel.CreatedAt,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(userModel.ID, userModel.Username, userModel.Email, userModel.Password, userModel.CreatedAt)
	return userDomain, nil
}
