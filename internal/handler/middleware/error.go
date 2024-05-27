package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"github.com/yinebebt/priceestimation/internal/constants/model/response"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	debugMode := viper.GetBool("debug")
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			e := c.Errors[0]
			err := e.Unwrap()

			resp := CastErrorResponse(err)
			if resp != nil {
				er := errorx.Cast(err)
				if debugMode {
					resp.Description = fmt.Sprintf("Error: %v", er)
					resp.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
				}
				response.SendErrorResponse(c, resp)
				return
			}
			response.SendErrorResponse(c, &response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Unknown server error",
			})
			return
		}
	}
}

func CastErrorResponse(err error) *response.ErrorResponse {
	for _, e := range errors.Error {
		if errorx.IsOfType(err, e.ErrorType) {
			er := errorx.Cast(err)
			resp := response.ErrorResponse{
				Code:       e.StatusCode,
				Message:    er.Message(),
				FieldError: response.ErrorFields(er.Cause()),
			}
			return &resp
		}
	}
	return nil
}
