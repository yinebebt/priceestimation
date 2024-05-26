package estimation

import (
	"github.com/gin-gonic/gin"
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
	"github.com/yinebebt/priceestimation/internal/constants/model/response"
	"github.com/yinebebt/priceestimation/internal/handler/rest"
	"github.com/yinebebt/priceestimation/internal/module"
	"github.com/yinebebt/priceestimation/platform/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type user struct {
	logger         logger.Logger
	module         module.User
	contextTimeout time.Duration
}

func InitUser(logger logger.Logger, module module.User, contextTimeout time.Duration) rest.User {
	return &user{
		module:         module,
		logger:         logger,
		contextTimeout: contextTimeout,
	}
}

// CreateUser
// @Summary      add user
// @Description  register user
// @Tags         Users
// @param 		 createUser body dto.User true "create user request body"
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.User
// @Failure      400  {object}  map[string]string
// @Failure		 403  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users [post]
func (u *user) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.User
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if err := req.Validate(); err != nil {
				err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
				u.logger.Info(ctx, "unable to bind user data", zap.Error(err))
				_ = ctx.Error(err)
				return
			}
		}
		createdUser, err := u.module.CreateUser(ctx, req)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		createdUser.Password = ""
		response.SuccessResponseData(ctx, http.StatusOK, createdUser)
	}
}

// GetUser
// @Summary      get user
// @Description  get user (sysAdmin)
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/:id [get]
func (u *user) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", ctx.Param("id"))
		usr, err := u.module.GetUser(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, usr)
	}
}

// DeleteUser godoc
// @Summary      delete user
// @Description  remove user
// @Tags         Brands
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/:id [delete]
func (u *user) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", ctx.Param("id"))
		err := u.module.DeleteUser(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, "User deleted")
	}
}

// LoginUser
// @Summary      login user
// @Description  login user
// @Tags         Users
// @param 		 Login body dto.LoginRequest true "login request body"
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.LogInResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/auth [post]
func (u *user) LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
			u.logger.Warn(ctx, "unable to bind user login data", zap.Error(err))
			_ = ctx.Error(err)
			return
		}
		usr, err := u.module.Login(ctx, req)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, usr)
	}
}
