package handlers

import (
	"net/http"

	"github.com/idmaksim/url-shortener-api/internal/errors"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	var (
		code                = errors.ErrCodeInternal
		message             = "Internal server error"
		status              = http.StatusInternalServerError
		details interface{} = nil
	)

	switch e := err.(type) {
	case *errors.HttpError:
		code, message, status, details = handleHttpError(e)
	case *errors.ServiceError:
		code, message, status, details = handleServiceError(e)
	case *echo.HTTPError:
		code, message, status = handleEchoError(e)
	}

	c.JSON(status, errors.ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	})
}

func handleHttpError(e *errors.HttpError) (string, string, int, interface{}) {
	return e.Code, e.Message, e.StatusCode, e.Details
}

func handleServiceError(e *errors.ServiceError) (string, string, int, interface{}) {
	status := http.StatusInternalServerError

	switch e.Code {
	case errors.ErrCodeNotFound:
		status = http.StatusNotFound
	case errors.ErrCodeInvalidCredentials:
		status = http.StatusUnauthorized
	case errors.ErrCodeConflict:
		status = http.StatusConflict
	}

	return e.Code, e.Message, status, e.Details
}

func handleEchoError(e *echo.HTTPError) (string, string, int) {
	code := errors.ErrCodeInternal
	status := e.Code
	message := e.Message.(string)

	switch status {
	case http.StatusNotFound:
		code = errors.ErrCodeNotFound
	case http.StatusBadRequest:
		code = errors.ErrCodeInvalidRequest
	case http.StatusUnauthorized:
		code = errors.ErrCodeUnauthorized
	}

	return code, message, status
}
