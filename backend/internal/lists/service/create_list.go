package lists_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *ListServiceImpl) CreateList(ctx context.Context, userID uuid.UUID, input CreateListInput) (domain.List, error) {
	if err := s.accessService.CanAccessBoardByID(ctx, userID, input.BoardID); err != nil {
		return domain.List{}, fmt.Errorf("cannot access board: %w", err)
	}

	lists, err := s.repo.GetBoardLists(ctx, input.BoardID)
	if err != nil {
		return domain.List{}, fmt.Errorf("get board lists: %w", err)
	}

	if len(lists) > 15 {
		return domain.List{}, fmt.Errorf("too many lists (%d), maximum=%d: %w", len(lists), 15, err)
	}

	var position int
	if len(lists) != 0 {
		lastList := lists[len(lists)-1]
		position = lastList.Position + 1
	}

	list := domain.NewListUninitialized(input.Title, input.BoardID, position)
	if err := list.Validate(); err != nil {
		return domain.List{}, fmt.Errorf("validate list: %w", err)
	}

	listDomain, err := s.repo.CreateList(ctx, list)
	if err != nil {
		return domain.List{}, fmt.Errorf("create list: %w", err)
	}

	return listDomain, nil
}
