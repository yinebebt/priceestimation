package port

import (
	"github.com/yinebebt/priceestimation/internal/constants/model/dto"
)

type PriceEstimationRepository interface {
	Create(estimation dto.PriceEstimation)
	Update(estimation dto.PriceEstimation)
	Delete(estimation dto.PriceEstimation)
	FindAll() []dto.PriceEstimation
}
