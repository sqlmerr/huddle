package lists_service

import (
	"context"

	"github.com/google/uuid"
	core_access "github.com/sqlmerr/huddle/backend/internal/core/access"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	lists_postgres_repository "github.com/sqlmerr/huddle/backend/internal/lists/repository/postgres"
)

type ListService interface {
	CreateList(ctx context.Context, userID uuid.UUID, input CreateListInput) (domain.List, error)
	GetList(ctx context.Context, userID uuid.UUID, listID uuid.UUID) (domain.List, error)
	GetBoardLists(ctx context.Context, userID uuid.UUID, boardID uuid.UUID) ([]domain.List, error)
	PatchList(ctx context.Context, userID uuid.UUID, input PatchListInput) (domain.List, error)
	DeleteList(ctx context.Context, userID uuid.UUID, listID uuid.UUID) error

	// TODO: ChangeListOrder
}

type ListServiceImpl struct {
	repo          lists_postgres_repository.ListRepository
	accessService core_access.AccessService
}

func NewListService(
	repo lists_postgres_repository.ListRepository,
	accessService core_access.AccessService,
) *ListServiceImpl {
	return &ListServiceImpl{repo, accessService}
}
