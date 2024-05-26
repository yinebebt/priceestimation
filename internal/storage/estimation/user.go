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

type user struct {
	db  persist.DB
	log logger.Logger
}

func InitUser(db persist.DB, log logger.Logger) storage.User {
	return &user{
		db:  db,
		log: log,
	}
}

func (u *user) CreateUser(ctx context.Context, param dto.User) (*dto.User, error) {
	userData, err := u.db.AddUser(ctx, db.AddUserParams{
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Email:     param.Email,
		Password:  param.Password,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not create user")
		u.log.Error(ctx, "unable to create user", zap.Error(err), zap.Any("user", param))
		return nil, err
	}
	return &dto.User{
		ID:        userData.ID,
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Password:  userData.Password,
		CreatedAt: userData.CreatedAt,
	}, nil
}

func (u *user) GetUser(ctx context.Context, id uuid.UUID) (*dto.User, error) {
	user, err := u.db.GetUser(ctx, id)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to get user", zap.Error(err))
		return nil, err
	}
	return &dto.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		Password:  user.Password,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *user) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := u.db.DeleteUser(ctx, id)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete user")
		u.log.Error(ctx, "unable to delete user", zap.Error(err))
		return err
	}
	return nil
}

func (u *user) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	user, err := u.db.GetUserByEmail(ctx, email)
	if err != nil {
		err = errors.ErrReadError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to get user", zap.Error(err))
		return nil, err
	}
	return &dto.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		Password:  user.Password,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
