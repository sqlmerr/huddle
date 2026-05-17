package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type Task struct {
	ID          uuid.UUID
	ListID      uuid.UUID
	Title       string
	Description *string
	Status      string
	Position    int
	CreatedAt   time.Time
}

func NewTask(
	id uuid.UUID,
	listID uuid.UUID,
	title string,
	description *string,
	status string,
	position int,
	createdAt time.Time,
) Task {
	return Task{
		ID:          id,
		ListID:      listID,
		Title:       title,
		Description: description,
		Status:      status,
		Position:    position,
		CreatedAt:   createdAt,
	}
}

func NewTaskUninitialized(
	listID uuid.UUID,
	title string,
	description *string,
	status string,
	position int,
) Task {
	return Task{
		ID:          UninitializedID,
		ListID:      listID,
		Title:       title,
		Description: description,
		Status:      status,
		Position:    position,
		CreatedAt:   UninitializedTime,
	}
}

func (t *Task) Validate() error {
	titleLength := len([]rune(t.Title))
	if titleLength == 0 || titleLength > 255 {
		return fmt.Errorf(
			"`Title` length must be between 1 and 255 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if t.Description != nil {
		descriptionLength := len([]rune(*t.Description))
		if descriptionLength == 0 || descriptionLength > 1000 {
			return fmt.Errorf(
				"`Description` length must be between 1 and 1000 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	statusLength := len([]rune(t.Status))
	if statusLength == 0 || statusLength > 50 {
		return fmt.Errorf(
			"`Status` length must be between 1 and 50 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
