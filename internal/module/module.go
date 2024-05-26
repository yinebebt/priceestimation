package module

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
)

type User interface {
	CreateUser(ctx context.Context, param dto.User) (*dto.User, error)
	GetUser(ctx context.Context) (*dto.User, error)
	DeleteUser(ctx context.Context) error
}

type Location interface {
	CreateLocation(ctx context.Context, param dto.Location) (*dto.Location, error)
	GetLocation(ctx context.Context) (*dto.Location, error)
	DeleteLocation(ctx context.Context) error
}

type PriceEstimation interface {
	CreatePriceEstimation(ctx context.Context, param dto.PriceEstimation) (*dto.PriceEstimation, error)
	GetPriceEstimation(ctx context.Context) (*dto.PriceEstimation, error)
	UpdatePriceEstimation(ctx context.Context, price decimal.Decimal) (*dto.PriceEstimation, error)
	DeletePriceEstimation(ctx context.Context) error
}
