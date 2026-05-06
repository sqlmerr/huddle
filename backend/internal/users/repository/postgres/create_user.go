package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (r *UsersRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
	RETURNING id, username, email, password, created_at;`
	row := r.pool.QueryRow(ctx, query, user.Username, user.Email, user.Password)

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
