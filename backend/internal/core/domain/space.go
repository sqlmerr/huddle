package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type Space struct {
	ID          uuid.UUID
	Title       string
	Description *string

	OwnerID uuid.UUID

	CreatedAt  time.Time
	IsArchived bool
}

func NewSpace(
	id uuid.UUID,
	title string,
	description *string,
	ownerID uuid.UUID,
	createdAt time.Time,
	isArchived bool,
) Space {
	return Space{
		ID:          id,
		Title:       title,
		Description: description,
		OwnerID:     ownerID,
		CreatedAt:   createdAt,
		IsArchived:  isArchived,
	}
}

func NewSpaceUninitialized(title string, description *string, ownerID uuid.UUID) Space {
	return Space{
		ID:          UninitializedID,
		Title:       title,
		Description: description,
		OwnerID:     ownerID,
		CreatedAt:   UninitializedTime,
		IsArchived:  UninitializedIsArchived,
	}
}

func (s *Space) Validate() error {
	titleLen := len([]rune(s.Title))
	if titleLen == 0 || titleLen > 50 {
		return fmt.Errorf(
			"length of the title should be from 1 to 50 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if s.Description != nil {
		descriptionLen := len([]rune(*s.Description))
		if descriptionLen == 0 || descriptionLen > 1000 {
			return fmt.Errorf(
				"length of the description should be from 1 to 1000 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}
