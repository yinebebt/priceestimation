package initiator

import (
	"github.com/yinebebt/priceestimation/internal/constants/query/persist"
	"github.com/yinebebt/priceestimation/internal/storage"
	"github.com/yinebebt/priceestimation/internal/storage/estimation"
	"github.com/yinebebt/priceestimation/platform/logger"
)

type Persistence struct {
	user            storage.User
	priceEstimation storage.PriceEstimation
	location        storage.Location
}

func InitPersistence(db persist.DB, log logger.Logger) Persistence {
	return Persistence{
		user:            estimation.InitUser(db, log.Named("user-persistence")),
		priceEstimation: estimation.InitPriceEstimation(db, log.Named("priceEstimation-persistence")),
		location:        estimation.InitLocation(db, log.Named("location-persistence")),
	}
}
