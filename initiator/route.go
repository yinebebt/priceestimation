package initiator

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yinebebt/priceestimation/docs"
	"github.com/yinebebt/priceestimation/internal/glue/routing/estimation"
	"github.com/yinebebt/priceestimation/internal/glue/routing/location"
	"github.com/yinebebt/priceestimation/internal/glue/routing/user"
	"github.com/yinebebt/priceestimation/platform/logger"
)

func InitRouter(group *gin.RouterGroup, handler Handler, log logger.Logger) {
	docs.SwaggerInfo.Host = viper.GetString("server.host")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = viper.GetStringSlice("swagger.schemes")
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.InitRoute(group, handler.user)
	location.InitRoute(group, handler.location)
	estimation.InitRoute(group, handler.priceEstimation)
}
