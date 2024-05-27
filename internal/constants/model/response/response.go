package response

import (
	"encoding/json"
	error2 "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
	"github.com/yinebebt/priceestimation/internal/constants/errors"
	"log"

	"net/http"
)

type ResponseObjectAL any

type Response struct {
	// OK is only true if the request was successful.
	OK bool `json:"ok"`
	// Data contains the actual data of the response.
	Data interface{} `json:"data,omitempty"`
	// Error contains the error detail if the request was not successful.
	Error *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	// Code is the error code. It is not status code
	Code int `json:"code"`
	// Message is the error message.
	Message string `json:"message,omitempty"`
	// Description is the error description.
	Description string `json:"description,omitempty"`
	// StackTrace is the stack trace of the error.
	// It is only returned for debugging
	StackTrace string `json:"stack_trace,omitempty"`
	// FieldError is the error detail for each field, if available that is.
	FieldError []FieldError `json:"field_error,omitempty"`
}

type FieldError struct {
	// Name is the name of the field that caused the error.
	Name string `json:"name"`
	// Description is the error description for this field.
	Description string `json:"description"`
}

func SuccessResponse(ctx *gin.Context, statusCode int, data interface{}) {
	obj, _ := json.Marshal(data) //nolint:errcheck // since a rare error can safely be ignored
	ctx.Set("response_object", ResponseObjectAL(string(obj)))
	ctx.Set("status_code", statusCode)

	ctx.JSON(
		statusCode,
		Response{
			OK:   true,
			Data: data,
		},
	)
}

func SuccessResponseData(ctx *gin.Context, statusCode int, data interface{}) {
	obj, err := json.Marshal(data)
	if err != nil {
		log.Printf("error while marshaling data: %v", err)
		return
	}
	ctx.Set("response_object", ResponseObjectAL(string(obj)))
	ctx.Set("status_code", statusCode)

	ctx.JSON(
		statusCode,
		Response{
			OK:   true,
			Data: data,
		},
	)
}

func SendErrorResponse(ctx *gin.Context, err *ErrorResponse) {
	ctx.Set("status_code", err.Code)
	ctx.AbortWithStatusJSON(err.Code, Response{
		OK:    false,
		Error: err,
	})
}

func GetErrorFrom(err error) *ErrorResponse {
	debugMode := viper.GetBool("debug")

	for _, e := range errors.Error {
		if errorx.IsOfType(err, e.ErrorType) {
			er := errorx.Cast(err)
			res := ErrorResponse{
				Code:       e.StatusCode,
				Message:    er.Message(),
				FieldError: ErrorFields(er.Cause()),
			}
			if debugMode {
				res.Description = fmt.Sprintf("Error: %v", er)
				res.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
			}
			return &res
		}
	}
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Unknown server error",
	}
}

func ErrorFields(err error) []FieldError {
	var errs []FieldError
	var data validation.Errors
	if error2.As(err, &data) {
		for i, v := range data {
			errs = append(errs, FieldError{
				Name:        i,
				Description: v.Error(),
			},
			)
		}
		return errs
	}
	return nil
}
