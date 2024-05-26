package initiator

import (
	"github.com/yinebebt/priceestimation/internal/module"
	"github.com/yinebebt/priceestimation/internal/module/estimation"
	"github.com/yinebebt/priceestimation/platform/logger"
)

type Module struct {
	user            module.User
	location        module.Location
	priceEstimation module.PriceEstimation
}

func InitModule(persistence Persistence, log logger.Logger) Module {
	return Module{
		user:            estimation.InitUser(log.Named("user-module"), persistence.user),
		location:        estimation.InitLocation(log.Named("location-module"), persistence.location),
		priceEstimation: estimation.InitPriceEstimation(log.Named("asset-module"), persistence.priceEstimation),
	}
}
