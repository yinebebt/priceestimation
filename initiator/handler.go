package initiator

import (
	"github.com/yinebebt/priceestimation/internal/handler/rest"
	"github.com/yinebebt/priceestimation/internal/handler/rest/estimation"
	"github.com/yinebebt/priceestimation/platform/logger"
	"time"
)

type Handler struct {
	user            rest.User
	location        rest.Location
	priceEstimation rest.PriceEstimation
}

func InitHandler(module Module, log logger.Logger, timeout time.Duration) Handler {
	return Handler{
		user:            estimation.InitUser(log.Named("user-handler"), module.user, timeout),
		location:        estimation.InitLocation(log.Named("location-handler"), module.location, timeout),
		priceEstimation: estimation.InitPriceEstimation(log.Named("price-estimation-handler"), module.priceEstimation, timeout),
	}
}
