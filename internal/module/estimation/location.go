package estimation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
	"github.com/yinebebt/priceestimation/internal/module"
	"github.com/yinebebt/priceestimation/internal/storage"
	"github.com/yinebebt/priceestimation/platform/logger"
	"go.uber.org/zap"
)

type location struct {
	storage storage.Location
	log     logger.Logger
}

func InitLocation(log logger.Logger, store storage.Location) module.Location {
	return &location{
		storage: store,
		log:     log,
	}
}

func (l *location) CreateLocation(ctx context.Context, param dto.Location) (*dto.Location, error) {
	var err error
	if err = param.Validate(); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		l.log.Info(ctx, "invalid input", zap.Error(err), zap.Any("input", param))
		return nil, err
	}
	return l.storage.CreateLocation(ctx, param)
}

func (l *location) GetLocation(ctx context.Context) (*dto.Location, error) {
	userID, err := getLocationID(ctx, l.log)
	if err != nil {
		return nil, err
	}
	return l.storage.GetLocation(ctx, userID)
}

func (l *location) DeleteLocation(ctx context.Context) error {
	userID, err := getLocationID(ctx, l.log)
	if err != nil {
		return err
	}
	usr, err := l.storage.GetLocation(ctx, userID)
	if err != nil {
		return err
	}
	return l.storage.DeleteLocation(ctx, usr.ID)
}

func getLocationID(ctx context.Context, log logger.Logger) (uuid.UUID, error) {
	locationID := ctx.Value("id")
	if locationID == nil {
		err := errors.ErrInvalidUserInput.Wrap(nil, "invalid input", "empty query param")
		log.Info(ctx, "unable to bind user data", zap.Error(err))
		return uuid.Nil, err
	}
	locationUID, err := uuid.Parse(fmt.Sprint(locationID))
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid uuid", fmt.Sprintf("id:%s", locationID))
		log.Warn(ctx, "invalid id", zap.Error(err))
		return uuid.Nil, err
	}
	return locationUID, nil
}
