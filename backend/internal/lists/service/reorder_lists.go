package lists_service

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *ListServiceImpl) ReorderLists(ctx context.Context, userID uuid.UUID, input ReorderListsInput) error {
	if err := s.accessService.CanAccessBoardByID(ctx, userID, input.BoardID); err != nil {
		return fmt.Errorf("cannot access board: %w", err)
	}

	lists, err := s.repo.GetBoardLists(ctx, input.BoardID)
	if err != nil {
		return fmt.Errorf("get board lists: %w", err)
	}

	if len(input.Order) != len(lists) {
		return fmt.Errorf(
			"`Order` must contain ALL list ids ONLY for this board: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	newLists := make([]domain.List, len(lists))

	for i, list := range lists {
		if !slices.Contains(input.Order, list.ID) {
			return fmt.Errorf(
				"`Order` must contain ALL list ids for this board: %w",
				core_errors.ErrInvalidArgument,
			)
		}
		newList := list
		newList.Position = slices.Index(input.Order, list.ID)
		newLists[i] = newList
	}

	for _, list := range newLists {
		_, err := s.repo.SaveList(ctx, list)
		if err != nil {
			return fmt.Errorf("save list: %w", err)
		}
	}

	return nil
}
