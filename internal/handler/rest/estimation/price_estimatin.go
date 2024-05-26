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

type priceEstimation struct {
	logger         logger.Logger
	module         module.PriceEstimation
	contextTimeout time.Duration
}

func InitPriceEstimation(logger logger.Logger, module module.PriceEstimation, contextTimeout time.Duration) rest.PriceEstimation {
	return &priceEstimation{
		module:         module,
		logger:         logger,
		contextTimeout: contextTimeout,
	}
}

// CreatePriceEstimation
// @Summary      add priceEstimation
// @Description  register priceEstimation
// @Tags         PriceEstimations
// @param 		 createPriceEstimation body dto.PriceEstimation true "create priceEstimation request body"
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.PriceEstimation
// @Failure      400  {object}  map[string]string
// @Failure		 403  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /estimations [post]
func (u *priceEstimation) CreatePriceEstimation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.PriceEstimation
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if err := req.Validate(); err != nil {
				err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
				u.logger.Info(ctx, "unable to bind priceEstimation data", zap.Error(err))
				_ = ctx.Error(err)
				return
			}
		}
		createdPriceEstimation, err := u.module.CreatePriceEstimation(ctx, req)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, createdPriceEstimation)
	}
}

// GetPriceEstimation
// @Summary      get priceEstimation
// @Description  get priceEstimation (sysAdmin)
// @Tags         PriceEstimations
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.PriceEstimation
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /estimations/:id [get]
func (u *priceEstimation) GetPriceEstimation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", ctx.Param("id"))
		usr, err := u.module.GetPriceEstimation(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, usr)
	}
}

// DeletePriceEstimation godoc
// @Summary      delete priceEstimation
// @Description  remove priceEstimation
// @Tags         Brands
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /estimations/:id [delete]
func (u *priceEstimation) DeletePriceEstimation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", ctx.Param("id"))
		err := u.module.DeletePriceEstimation(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, "PriceEstimation deleted")
	}
}

// UpdatePriceEstimation
// @Summary      update priceEstimation
// @Description  update priceEstimation
// @Tags         PriceEstimations
// @param 		 createPriceEstimation body dto.PriceEstimation true "create priceEstimation request body"
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.PriceEstimation
// @Failure      400  {object}  map[string]string
// @Failure		 403  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /estimations/:id [patch]
func (u *priceEstimation) UpdatePriceEstimation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.UpdatePriceEstimation
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if err := req.Validate(); err != nil {
				err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
				u.logger.Info(ctx, "unable to bind priceEstimation data", zap.Error(err))
				_ = ctx.Error(err)
				return
			}
		}
		updatedPriceEstimation, err := u.module.UpdatePriceEstimation(ctx, req.Price)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, updatedPriceEstimation)
	}
}

// ListPriceEstimation
// @Summary      get priceEstimation
// @Description  get priceEstimation (sysAdmin)
// @Tags         PriceEstimations
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.PaginationRequest
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /estimations [get]
func (u *priceEstimation) ListPriceEstimation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pagination dto.PaginationRequest
		if err := ctx.ShouldBindQuery(&pagination); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		est, err := u.module.ListPriceEstimation(ctx, pagination)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		response.SuccessResponseData(ctx, http.StatusOK, est)
	}
}
