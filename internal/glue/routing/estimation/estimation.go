package estimation

import (
	"github.com/yinebebt/priceestimation/internal/glue/routing"
	"github.com/yinebebt/priceestimation/internal/handler/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute(grp *gin.RouterGroup, estimation rest.PriceEstimation) {
	userRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Path:    "",
			Handler: estimation.CreatePriceEstimation(),
		},
		{
			Method:  http.MethodGet,
			Path:    "/:id",
			Handler: estimation.GetPriceEstimation(),
		},
		{
			Method:  http.MethodDelete,
			Path:    "/:id",
			Handler: estimation.DeletePriceEstimation(),
		},
		{
			Method:  http.MethodPatch,
			Path:    "/:id",
			Handler: estimation.UpdatePriceEstimation(),
		},
		{
			Method:  http.MethodGet,
			Path:    "",
			Handler: estimation.ListPriceEstimation(),
		},
	}

	routing.RegisterRoutes(grp.Group("/estimations"), userRoutes)
}
