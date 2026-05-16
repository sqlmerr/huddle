package core_access

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type AccessService interface {
	CanAccessSpaceByID(ctx context.Context, userID, spaceID uuid.UUID) error
	CanAccessSpace(ctx context.Context, userID uuid.UUID, space domain.Space) error
	CanAccessBoardByID(ctx context.Context, userID, boardID uuid.UUID) error
	CanAccessBoard(ctx context.Context, userID uuid.UUID, board domain.Board) error
	CanAccessList(ctx context.Context, userID uuid.UUID, list domain.List) error
	CanAccessListByID(ctx context.Context, userID, listID uuid.UUID) error
	CanAccessTask(ctx context.Context, userID uuid.UUID, task domain.Task) error
	CanAccessTaskByID(ctx context.Context, userID, taskID uuid.UUID) error
}

type SpaceRepository interface {
	GetSpace(ctx context.Context, spaceID uuid.UUID) (domain.Space, error)
}

type BoardRepository interface {
	GetBoard(ctx context.Context, boardID uuid.UUID) (domain.Board, error)
}

type ListRepository interface {
	GetList(ctx context.Context, listID uuid.UUID) (domain.List, error)
}

type TaskRepository interface {
	GetTask(ctx context.Context, taskID uuid.UUID) (domain.Task, error)
}
