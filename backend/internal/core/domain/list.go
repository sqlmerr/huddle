package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type List struct {
	ID        uuid.UUID
	Title     string
	BoardID   uuid.UUID
	Position  int
	CreatedAt time.Time
}

func NewList(
	id uuid.UUID,
	title string,
	boardID uuid.UUID,
	position int,
	createdAt time.Time,
) List {
	return List{
		id,
		title,
		boardID,
		position,
		createdAt,
	}
}

func NewListUninitialized(
	title string,
	boardID uuid.UUID,
	position int,
) List {
	return List{
		ID:        UninitializedID,
		Title:     title,
		BoardID:   boardID,
		Position:  position,
		CreatedAt: UninitializedTime,
	}
}

func (l *List) Validate() error {
	titleLen := len([]rune(l.Title))
	if titleLen == 0 || titleLen > 50 {
		return fmt.Errorf("`Title` length must be between 1 and 50 characters: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
