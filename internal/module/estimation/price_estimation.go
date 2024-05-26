package estimation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
	"github.com/yinebebt/priceestimation/internal/module"
	"github.com/yinebebt/priceestimation/internal/storage"
	"github.com/yinebebt/priceestimation/platform/logger"
	"go.uber.org/zap"
)

type priceEstimation struct {
	storage storage.PriceEstimation
	log     logger.Logger
}

func InitPriceEstimation(log logger.Logger, store storage.PriceEstimation) module.PriceEstimation {
	return &priceEstimation{
		storage: store,
		log:     log,
	}
}

func (p *priceEstimation) CreatePriceEstimation(ctx context.Context, param dto.PriceEstimation) (*dto.PriceEstimation, error) {
	var err error
	if err = param.Validate(); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.log.Info(ctx, "invalid input", zap.Error(err), zap.Any("input", param))
		return nil, err
	}
	return p.storage.CreatePriceEstimation(ctx, param)
}

func (p *priceEstimation) GetPriceEstimation(ctx context.Context) (*dto.PriceEstimation, error) {
	estID, err := getPriceEstimationID(ctx, p.log)
	if err != nil {
		return nil, err
	}
	return p.storage.GetPriceEstimation(ctx, estID)
}

func (p *priceEstimation) DeletePriceEstimation(ctx context.Context) error {
	userID, err := getPriceEstimationID(ctx, p.log)
	if err != nil {
		return err
	}
	est, err := p.storage.GetPriceEstimation(ctx, userID)
	if err != nil {
		return err
	}
	return p.storage.DeletePriceEstimation(ctx, est.ID)
}

func (p *priceEstimation) UpdatePriceEstimation(ctx context.Context, price decimal.Decimal) (*dto.PriceEstimation, error) {
	return p.storage.UpdatePriceEstimation(ctx, price)
}

func getPriceEstimationID(ctx context.Context, log logger.Logger) (uuid.UUID, error) {
	userID := ctx.Value("id")
	if userID == nil {
		err := errors.ErrInvalidUserInput.Wrap(nil, "invalid input", "empty query param")
		log.Info(ctx, "unable to bind price estimation data", zap.Error(err))
		return uuid.Nil, err
	}
	userUID, err := uuid.Parse(fmt.Sprint(userID))
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid uuid", fmt.Sprintf("id:%s", userID))
		log.Warn(ctx, "invalid id", zap.Error(err))
		return uuid.Nil, err
	}
	return userUID, nil
}

func (p *priceEstimation) ListPriceEstimation(ctx context.Context, param dto.PaginationRequest) ([]dto.PriceEstimation, error) {
	// do validation on pagination data
	return p.storage.ListPriceEstimation(ctx, param)
}
