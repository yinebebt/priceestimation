package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/yinebebt/priceestimation/platform/logger"
	"github.com/yinebebt/priceestimation/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strings"
	"time"
)

// InitCORS is used to set cross-origin resource sharing policies
func InitCORS() gin.HandlerFunc {
	origins := viper.GetStringSlice("cors.origin")
	if len(origins) == 0 {
		origins = []string{"*"}
	}
	allowCredentials := viper.GetString("cors.allow_credentials")
	if allowCredentials == "" {
		allowCredentials = "true"
	}
	headers := viper.GetStringSlice("cors.headers")
	if len(headers) == 0 {
		headers = []string{"*"}
	}
	methods := viper.GetStringSlice("cors.methods")
	if len(methods) == 0 {
		methods = []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"PATCH",
			"OPTIONS",
		}
	}

	return func(c *gin.Context) {
		currentOrigin := c.Request.Header.Get("Origin")
		if utils.Contains(currentOrigin, origins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", currentOrigin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origins[0])
		}
		// Set the necessary headers to allow CORS
		c.Writer.Header().Set("Access-Control-Allow-Credentials", allowCredentials)
		c.Header("Access-Control-Allow-Methods", strings.Join(methods, ","))
		c.Header("Access-Control-Allow-Headers", strings.Join(headers, ","))
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// GinLogger is used to log gin-specific log fields
func GinLogger(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		id := uuid.New().String()
		ctx.Set("x-request-id", id)
		ctx.Set("request-start-time", start)
		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)
		fields := []zapcore.Field{
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.Int64("request-latency", latency.Milliseconds()),
		}
		log.Info(ctx, "GIN", fields...)
	}
}
