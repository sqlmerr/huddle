package auth_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *AuthRepositoryImpl) GetUser(ctx context.Context, userID uuid.UUID) (domain.User, error) {
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

func (r *AuthRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, username, email, password, created_at
	FROM users
	WHERE lower(username) = $1
	`
	row := r.pool.QueryRow(ctx, query, strings.ToLower(username))

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Username,
		&userModel.Email,
		&userModel.Password,
		&userModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with username '%s': %w", username, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(userModel.ID, userModel.Username, userModel.Email, userModel.Password, userModel.CreatedAt)
	return userDomain, nil
}

func (r *AuthRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, username, email, password, created_at
	FROM users
	WHERE lower(email) = $1
	`
	row := r.pool.QueryRow(ctx, query, strings.ToLower(email))

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Username,
		&userModel.Email,
		&userModel.Password,
		&userModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with email '%s': %w", email, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(userModel.ID, userModel.Username, userModel.Email, userModel.Password, userModel.CreatedAt)
	return userDomain, nil
}
