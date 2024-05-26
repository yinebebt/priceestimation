package initiator

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yinebebt/priceestimation/internal/glue/routing/user"
	"github.com/yinebebt/priceestimation/platform/logger"
)

func InitRouter(group *gin.RouterGroup, handler Handler, log logger.Logger) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.InitRoute(group, handler.user)
}
