// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Querier interface {
	AddLocation(ctx context.Context, arg AddLocationParams) (Location, error)
	AddPriceEstimation(ctx context.Context, arg AddPriceEstimationParams) (PriceEstimation, error)
	AddUser(ctx context.Context, arg AddUserParams) (User, error)
	DeleteLocation(ctx context.Context, id uuid.UUID) error
	DeletePriceEstimation(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetLocation(ctx context.Context, id uuid.UUID) (Location, error)
	GetLocations(ctx context.Context) ([]Location, error)
	GetPriceEstimation(ctx context.Context, id uuid.UUID) (GetPriceEstimationRow, error)
	GetPriceEstimations(ctx context.Context) ([]PriceEstimation, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	UpdatePriceEstimation(ctx context.Context, price decimal.Decimal) (PriceEstimation, error)
}

var _ Querier = (*Queries)(nil)
