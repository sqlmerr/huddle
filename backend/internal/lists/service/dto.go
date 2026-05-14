package lists_service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type CreateListInput struct {
	Title   string
	BoardID uuid.UUID
}

type PatchListInput struct {
	ListID uuid.UUID
	Title  domain.Nullable[string]
}

func (i *PatchListInput) Validate() error {
	if i.Title.Set && i.Title.Value == nil {
		return fmt.Errorf("`Title` can't be nil")
	}

	return nil
}

func (i *PatchListInput) ApplyPatch(list domain.List) (domain.List, error) {
	if err := i.Validate(); err != nil {
		return domain.List{}, fmt.Errorf("validate patch list input: %w", err)
	}

	tmp := list
	if i.Title.Set {
		tmp.Title = *i.Title.Value
	}

	if err := tmp.Validate(); err != nil {
		return domain.List{}, fmt.Errorf("validate patch list input: %w", err)
	}

	return tmp, nil
}
