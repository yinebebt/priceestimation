package user

import (
	"github.com/yinebebt/priceestimation/internal/glue/routing"
	"github.com/yinebebt/priceestimation/internal/handler/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute(grp *gin.RouterGroup, user rest.User) {
	userRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Path:    "",
			Handler: user.CreateUser(),
		},
		{
			Method:  http.MethodGet,
			Path:    "/:id",
			Handler: user.GetUser(),
		}, {
			Method:  http.MethodDelete,
			Path:    "/:id",
			Handler: user.DeleteUser(),
		},
	}

	routing.RegisterRoutes(grp.Group("/users"), userRoutes)
}
