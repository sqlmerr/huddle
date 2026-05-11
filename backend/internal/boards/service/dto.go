package boards_service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type CreateBoardInput struct {
	Title   string
	SpaceID uuid.UUID
}

type PatchBoardInput struct {
	BoardID uuid.UUID

	Title domain.Nullable[string]
}

func (i *PatchBoardInput) Validate() error {
	if i.Title.Set && i.Title.Value == nil {
		return fmt.Errorf("`Title` can't be nil")
	}

	return nil
}

func (i *PatchBoardInput) ApplyPatch(board domain.Board) (domain.Board, error) {
	if err := i.Validate(); err != nil {
		return domain.Board{}, fmt.Errorf("validate patch input: %w", err)
	}

	tmp := board

	if i.Title.Set {
		tmp.Title = *i.Title.Value
	}

	if err := tmp.Validate(); err != nil {
		return domain.Board{}, fmt.Errorf("validate patch input: %w", err)
	}
	return tmp, nil
}
