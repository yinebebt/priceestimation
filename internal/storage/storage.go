package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
)

type PriceEstimation interface {
	Create(ctx context.Context, param dto.User) (*dto.User, error)
	Get(ctx context.Context, id uuid.UUID) (*dto.User, error)
	Update(ctx context.Context, param dto.User) (*dto.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type Location interface {
	Create(ctx context.Context, param dto.Location) (*dto.Location, error)
	Get(ctx context.Context, id uuid.UUID) (*dto.Location, error)
	Update(ctx context.Context, param dto.Location) (*dto.Location, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type User interface {
	Create(ctx context.Context, param dto.User) (*dto.User, error)
	Get(ctx context.Context, id uuid.UUID) (*dto.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
