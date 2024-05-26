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
	"github.com/yinebebt/priceestimation/utils"
	"go.uber.org/zap"
)

type user struct {
	storage storage.User
	log     logger.Logger
}

func InitUser(log logger.Logger, store storage.User) module.User {
	return &user{
		storage: store,
		log:     log,
	}
}

func (u *user) CreateUser(ctx context.Context, param dto.User) (*dto.User, error) {
	var err error
	if err = param.Validate(); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.log.Info(ctx, "invalid input", zap.Error(err), zap.Any("input", param))
		return nil, err
	}
	hashedPassword, err := utils.HashPassword(param.Password)
	if err != nil {
		err := errors.ErrInternalServerError.Wrap(err, "user password hashing failed")
		u.log.Info(ctx, "invalid input", zap.Error(err), zap.Any("input", param))
		return nil, err
	}
	param.Password = hashedPassword
	return u.storage.CreateUser(ctx, param)
}

func (u *user) GetUser(ctx context.Context) (*dto.User, error) {
	userID, err := getUserID(ctx, u.log)
	if err != nil {
		return nil, err
	}
	return u.storage.GetUser(ctx, userID)
}

func (u *user) DeleteUser(ctx context.Context) error {
	userID, err := getUserID(ctx, u.log)
	if err != nil {
		return err
	}
	usr, err := u.storage.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	return u.storage.DeleteUser(ctx, usr.ID)
}

func getUserID(ctx context.Context, log logger.Logger) (uuid.UUID, error) {
	userID := ctx.Value("id")
	if userID == nil {
		err := errors.ErrInvalidUserInput.Wrap(nil, "invalid input", "empty query param")
		log.Info(ctx, "unable to bind user data", zap.Error(err))
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
