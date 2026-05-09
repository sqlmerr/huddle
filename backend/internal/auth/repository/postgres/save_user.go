package auth_postgres_repository

import (
	"context"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (r *AuthRepositoryImpl) SaveUser(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE users
		SET
			username = $2,
			email = $3,
			password = $4
		WHERE id = $1
	`

	commandTag, err := r.pool.Exec(ctx, query, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%s' not found: %w", user.ID.String(), core_errors.ErrNotFound)
	}

	return nil
}
