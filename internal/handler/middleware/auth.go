package middleware

import (
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"github.com/yinebebt/priceestimation/platform/logger"
	"github.com/yinebebt/priceestimation/utils"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Authenticate() gin.HandlerFunc
}

type authMiddleware struct {
	logger logger.Logger
}

func InitAuthMiddleware(
	logger logger.Logger) AuthMiddleware {
	return &authMiddleware{
		logger: logger,
	}
}

// Do authorization here
//

// Authenticate make authentication with JWT token
func (a *authMiddleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			a.logger.Warn(ctx, "Authorization header not found")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header not found"})
			return
		}
		tokenString := authHeader[len(BearerSchema):]
		claims, valid := utils.ValidateToken(tokenString)
		if !valid {
			err := errors.ErrInvalidAccessToken.New("Unauthorized")
			a.logger.Warn(ctx, "invalid token was tried", zap.String("token", tokenString), zap.Error(err))

			_ = ctx.Error(err)
			ctx.Abort()
			return
		}
		//check if the token is from known client,whether the token is not for us, scope and etc
		if len(claims.Audience) != 0 {
			if !utils.Contains("https://chipchip.social", claims.Audience) {
				//todo: read this via state, check scope
				err := errors.ErrInvalidAccessToken.New("Unauthorized")
				a.logger.Warn(ctx, "token didn't have the required scope",
					zap.Any("token", claims),
					zap.Error(err))

				_ = ctx.Error(err)
				ctx.Abort()
				return
			}
		} else {
			err := errors.ErrInvalidAccessToken.New("Unauthorized")
			a.logger.Warn(ctx, "token without an audience was used!",
				zap.Any("token", claims),
				zap.Error(err))

			_ = ctx.Error(err)
			ctx.Abort()

			return
		}
		//todo: check user status, do extra logics
		if err := claims.Valid(); err != nil {
			err := errors.ErrInvalidAccessToken.New("Unauthorized")
			a.logger.Warn(ctx, "expired(invalid) token was used!",
				zap.Any("token", claims),
				zap.Error(err))

			_ = ctx.Error(err)
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims.Subject)
		ctx.Next()
	}
}
