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

type location struct {
	logger         logger.Logger
	module         module.Location
	contextTimeout time.Duration
}

func InitLocation(logger logger.Logger, module module.Location, contextTimeout time.Duration) rest.Location {
	return &location{
		module:         module,
		logger:         logger,
		contextTimeout: contextTimeout,
	}
}

// CreateLocation
// @Summary      add location
// @Description  register location
// @Tags         Users
// @param 		 createUser body dto.Location true "create location request body"
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Location
// @Failure      400  {object}  map[string]string
// @Failure		 403  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /locations [post]
func (l *location) CreateLocation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.Location
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if err := req.Validate(); err != nil {
				err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
				l.logger.Info(ctx, "unable to bind location data", zap.Error(err))
				_ = ctx.Error(err)
				return
			}
		}
		createdLoc, err := l.module.CreateLocation(ctx, req)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, createdLoc)
	}
}

// GetLocation
// @Summary      get location
// @Description  get location
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Location
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /locations/:id [get]
func (l *location) GetLocation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", ctx.Param("id"))
		usr, err := l.module.GetLocation(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, usr)
	}
}
