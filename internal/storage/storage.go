package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
)

type PriceEstimation interface {
	CreatePriceEstimation(ctx context.Context, param dto.PriceEstimation) (*dto.PriceEstimation, error)
	GetPriceEstimation(ctx context.Context, id uuid.UUID) (*dto.PriceEstimation, error)
	UpdatePriceEstimation(ctx context.Context, price decimal.Decimal) (*dto.PriceEstimation, error)
	DeletePriceEstimation(ctx context.Context, id uuid.UUID) error
}

type Location interface {
	CreateLocation(ctx context.Context, param dto.Location) (*dto.Location, error)
	GetLocation(ctx context.Context, id uuid.UUID) (*dto.Location, error)
	DeleteLocation(ctx context.Context, id uuid.UUID) error
}

type User interface {
	CreateUser(ctx context.Context, param dto.User) (*dto.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*dto.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
