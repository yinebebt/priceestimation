package estimation

import (
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"github.com/yinebebt/priceestimation/internal/constants/model/db"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
	"github.com/yinebebt/priceestimation/internal/constants/query/persist"
	"github.com/yinebebt/priceestimation/internal/storage"
	"github.com/yinebebt/priceestimation/platform/logger"
	"go.uber.org/zap"
)

type priceEstimation struct {
	db  persist.DB
	log logger.Logger
}

func InitPriceEstimation(db persist.DB, log logger.Logger) storage.PriceEstimation {
	return &priceEstimation{
		db:  db,
		log: log,
	}
}

func (u *priceEstimation) CreatePriceEstimation(ctx context.Context, param dto.PriceEstimation) (*dto.PriceEstimation, error) {
	estData, err := u.db.AddPriceEstimation(ctx, db.AddPriceEstimationParams{
		ProductName: param.ProductName,
		Price:       param.Price,
		UserID:      param.UserID,
		LocationID:  param.LocationID,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not create user")
		u.log.Error(ctx, "unable to create user", zap.Error(err), zap.Any("user", param))
		return nil, err
	}
	return &dto.PriceEstimation{
		ID:          estData.ID,
		ProductName: estData.ProductName,
		Price:       estData.Price,
		UserID:      estData.UserID,
		LocationID:  estData.LocationID,
		CreatedAt:   estData.CreatedAt,
		UpdatedAt:   estData.UpdatedAt,
	}, nil
}

func (u *priceEstimation) GetPriceEstimation(ctx context.Context, id uuid.UUID) (*dto.PriceEstimation, error) {
	estData, err := u.db.GetPriceEstimation(ctx, id)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to get user", zap.Error(err))
		return nil, err
	}
	return &dto.PriceEstimation{
		ID:        estData.ID,
		Price:     estData.Price,
		UserID:    estData.UserID,
		CreatedAt: estData.CreatedAt,
		UpdatedAt: estData.UpdatedAt,
		Location: dto.Location{
			Country: estData.LocationCountry,
			Region:  estData.LocationRegion,
			Zone:    estData.LocationZone,
			City:    estData.LocationCity,
		},
	}, nil
}

func (u *priceEstimation) UpdatePriceEstimation(ctx context.Context, price decimal.Decimal) (*dto.PriceEstimation, error) {
	estPrice, err := u.db.UpdatePriceEstimation(ctx, price)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete price estimation")
		u.log.Error(ctx, "unable to delete price estimation", zap.Error(err))
		return nil, err
	}
	return &dto.PriceEstimation{
		ID:          estPrice.ID,
		ProductName: estPrice.ProductName,
		Price:       estPrice.Price,
		UserID:      estPrice.UserID,
		LocationID:  estPrice.LocationID,
		CreatedAt:   estPrice.CreatedAt,
		UpdatedAt:   estPrice.UpdatedAt,
	}, nil
}

func (u *priceEstimation) ListPriceEstimation(ctx context.Context, param dto.PaginationRequest) ([]dto.PriceEstimation, error) {
	estData, err := u.db.ListPriceEstimation(ctx, db.ListPriceEstimationParams{
		Limit:  param.Limit,
		Offset: param.Offset,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to get user", zap.Error(err))
		return nil, err
	}
	var est []dto.PriceEstimation
	for _, estimation := range estData {
		est = append(est, dto.PriceEstimation{
			ID:        estimation.ID,
			Price:     estimation.Price,
			UserID:    estimation.UserID,
			CreatedAt: estimation.CreatedAt,
			UpdatedAt: estimation.UpdatedAt,
			Location: dto.Location{
				Country: estimation.LocationCountry,
				Region:  estimation.LocationRegion,
				Zone:    estimation.LocationZone,
				City:    estimation.LocationCity,
			},
		})
	}
	return est, nil
}

func (u *priceEstimation) DeletePriceEstimation(ctx context.Context, id uuid.UUID) error {
	err := u.db.DeletePriceEstimation(ctx, id)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete price estimation")
		u.log.Error(ctx, "unable to delete price estimation", zap.Error(err))
		return err
	}
	return nil
}
