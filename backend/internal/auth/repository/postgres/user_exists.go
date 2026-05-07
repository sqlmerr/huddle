package auth_postgres_repository

import (
	"context"
	"strings"
)

func (r *AuthRepository) UserExistsByUsername(
	ctx context.Context,
	username string,
) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT EXISTS(SELECT 1 FROM users WHERE lower(username) = $1)
	`
	row := r.pool.QueryRow(ctx, query, strings.ToLower(username))

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *AuthRepository) UserExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT EXISTS(SELECT 1 FROM users WHERE lower(email) = $1)
	`
	row := r.pool.QueryRow(ctx, query, strings.ToLower(email))

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
