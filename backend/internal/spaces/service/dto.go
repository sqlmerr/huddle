package spaces_service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type CreateSpaceInput struct {
	OwnerID     uuid.UUID
	Title       string
	Description *string
}

type PatchSpaceInput struct {
	UserID      uuid.UUID
	SpaceID     uuid.UUID
	Title       domain.Nullable[string]
	Description domain.Nullable[string]
}

func (i *PatchSpaceInput) Validate() error {
	if i.Title.Set && i.Title.Value == nil {
		return fmt.Errorf("`Title` can't be null: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (i *PatchSpaceInput) ApplyPatch(d domain.Space) (domain.Space, error) {
	if err := i.Validate(); err != nil {
		return domain.Space{}, fmt.Errorf("validate patch request: %w", err)
	}

	tmp := d
	if i.Title.Set {
		tmp.Title = *i.Title.Value
	}
	if i.Description.Set {
		tmp.Description = i.Description.Value
	}

	if err := tmp.Validate(); err != nil {
		return domain.Space{}, fmt.Errorf("validate patch request: %w", err)
	}

	return tmp, nil
}
