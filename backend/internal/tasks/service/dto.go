package tasks_service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type CreateTaskInput struct {
	ListID      uuid.UUID
	Title       string
	Description *string
	Status      string
}

type ReorderTasksInput struct {
	ListID uuid.UUID
	Order  []uuid.UUID
}

type PatchTaskInput struct {
	TaskID uuid.UUID

	Title       domain.Nullable[string]
	Description domain.Nullable[string]
	Status      domain.Nullable[string]
}

func (i *PatchTaskInput) Validate() error {
	if i.Title.Set && i.Title.Value == nil {
		return fmt.Errorf("`Title` can't be nil: %w", core_errors.ErrInvalidArgument)
	}

	if i.Status.Set && i.Status.Value == nil {
		return fmt.Errorf("`Status` can't be nil: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (i *PatchTaskInput) ApplyPatch(task domain.Task) (domain.Task, error) {
	if err := i.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate patch task input: %w", err)
	}

	tmp := task
	if i.Title.Set {
		tmp.Title = *i.Title.Value
	}

	if i.Description.Set {
		tmp.Description = i.Description.Value
	}

	if i.Status.Set {
		tmp.Status = *i.Status.Value
	}

	if err := tmp.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate patch task input: %w", err)
	}

	return tmp, nil
}
