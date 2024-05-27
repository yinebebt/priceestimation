package location

import (
	"github.com/gin-gonic/gin"
	"github.com/yinebebt/priceestimation/internal/glue/routing"
	"github.com/yinebebt/priceestimation/internal/handler/rest"
	"net/http"
)

func InitRoute(grp *gin.RouterGroup, loc rest.Location) {
	userRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Path:    "",
			Handler: loc.CreateLocation(),
		},
		{
			Method:  http.MethodGet,
			Path:    "/:id",
			Handler: loc.GetLocation(),
		},
	}

	routing.RegisterRoutes(grp.Group("/locations"), userRoutes)
}
