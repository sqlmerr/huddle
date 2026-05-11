package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type Board struct {
	ID      uuid.UUID
	Title   string
	SpaceID uuid.UUID

	CreatedAt time.Time
}

func NewBoard(
	id uuid.UUID,
	title string,
	spaceID uuid.UUID,
	createdAt time.Time,
) Board {
	return Board{
		ID:        id,
		Title:     title,
		SpaceID:   spaceID,
		CreatedAt: createdAt,
	}
}

func NewBoardUninitialized(
	title string,
	spaceID uuid.UUID,
) Board {
	return Board{
		ID:        UninitializedID,
		Title:     title,
		SpaceID:   spaceID,
		CreatedAt: UninitializedTime,
	}
}

func (b *Board) Validate() error {
	titleLen := len([]rune(b.Title))
	if titleLen == 0 || titleLen > 50 {
		return fmt.Errorf(
			"length of the `Title` should be from 1 to 50 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
