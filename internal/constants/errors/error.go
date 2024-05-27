package errors

import (
	"github.com/joomcode/errorx"
	"net/http"
)

type ErrorType struct {
	StatusCode int
	ErrorType  *errorx.Type
}

var Error = []ErrorType{
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrInvalidUserInput,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrDataExists,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrReadError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrWriteError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrUpdateError,
	},
	{
		StatusCode: http.StatusNotFound,
		ErrorType:  ErrNoRecordFound,
	},
	{
		StatusCode: http.StatusUnauthorized,
		ErrorType:  ErrInvalidAccessToken,
	},
	{
		StatusCode: http.StatusForbidden,
		ErrorType:  ErrAccessError,
	},
}

// list of error namespaces
var (
	invalidInput     = errorx.NewNamespace("validation error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	dbError          = errorx.NewNamespace("db error")
	duplicate        = errorx.NewNamespace("duplicate").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	dataNotFound     = errorx.NewNamespace("data not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	serverError      = errorx.NewNamespace("server error")
	databaseError    = errorx.NewNamespace("database error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	resourceNotFound = errorx.NewNamespace("not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	Unauthenticated  = errorx.NewNamespace("user authentication failed")
	unauthorized     = errorx.NewNamespace("unauthorized").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	AccessDenied     = errorx.RegisterTrait("You are not authorized to perform the action")
)

// list of errors types in all of the above namespaces
var (
	ErrInvalidUserInput    = errorx.NewType(invalidInput, "invalid user input")
	ErrWriteError          = errorx.NewType(dbError, "could not write to db")
	ErrReadError           = errorx.NewType(dbError, "could not read data from db")
	ErrDataExists          = errorx.NewType(duplicate, "data already exists")
	ErrNoRecordFound       = errorx.NewType(dataNotFound, "no record found")
	ErrDeleteError         = errorx.NewType(dbError, "could not delete from db")
	ErrInternalServerError = errorx.NewType(serverError, "internal server error")
	ErrUnableToGet         = errorx.NewType(databaseError, "unable to get")
	ErrResourceNotFound    = errorx.NewType(resourceNotFound, "resource not found")
	ErrInvalidAccessToken  = errorx.NewType(Unauthenticated, "invalid token").
				ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	ErrAccessError = errorx.NewType(unauthorized, "Unauthorized", AccessDenied)
	ErrUpdateError = errorx.NewType(dbError, "could not update from db")
)
