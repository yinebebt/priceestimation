package estimation

import (
	"context"
	"github.com/google/uuid"
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"github.com/yinebebt/priceestimation/internal/constants/model/db"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
	"github.com/yinebebt/priceestimation/internal/constants/query/persist"
	"github.com/yinebebt/priceestimation/internal/storage"
	"github.com/yinebebt/priceestimation/platform/logger"
	"go.uber.org/zap"
)

type location struct {
	db  persist.DB
	log logger.Logger
}

func InitLocation(db persist.DB, log logger.Logger) storage.Location {
	return &location{
		db:  db,
		log: log,
	}
}

func (u *location) CreateLocation(ctx context.Context, param dto.Location) (*dto.Location, error) {
	locationData, err := u.db.AddLocation(ctx, db.AddLocationParams{
		Country: param.Country,
		Region:  param.Region,
		Zone:    param.Zone,
		City:    param.City,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not create location")
		u.log.Error(ctx, "unable to create location", zap.Error(err), zap.Any("location", param))
		return nil, err
	}
	return &dto.Location{
		ID:        locationData.ID,
		Country:   locationData.Country,
		Region:    locationData.Region,
		Zone:      locationData.Zone,
		City:      locationData.City,
		CreatedAt: locationData.CreatedAt,
	}, nil
}

func (u *location) GetLocation(ctx context.Context, id uuid.UUID) (*dto.Location, error) {
	location, err := u.db.GetLocation(ctx, id)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not read location")
		u.log.Error(ctx, "unable to get location", zap.Error(err))
		return nil, err
	}
	return &dto.Location{
		ID:        location.ID,
		Country:   location.Country,
		Region:    location.Region,
		Zone:      location.Zone,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}, nil
}

func (u *location) DeleteLocation(ctx context.Context, id uuid.UUID) error {
	err := u.db.DeleteLocation(ctx, id)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete location")
		u.log.Error(ctx, "unable to delete location", zap.Error(err))
		return err
	}
	return nil
}
