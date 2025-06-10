package sp

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"net/http"
)

func Builder() *Error {
	return New(Sample{
		Level: levels.LevelError,
	}).path(1)
}

func Any(caused error, desc, hint string) *Error {
	return formError(http.StatusInternalServerError, caused, "Internal Server Error", desc, hint).path(1)
}

func Internal(caused error, desc, hint string) *Error {
	return formError(http.StatusInternalServerError, caused, "Internal Server Error", desc, hint).path(1)
}

func NotFound(desc, hint string) *Error {
	return formError(http.StatusNotFound, nil, "Not Found", desc, hint).path(1)
}

func Forbidden(desc, hint string) *Error {
	return formError(http.StatusForbidden, nil, "Forbidden", desc, hint).path(1)
}

func BadRequest(desc, hint string) *Error {
	return formError(http.StatusBadRequest, nil, "Bad Request", desc, hint).path(1)
}

func formError(code int, err error, msg, desc, hint string) *Error {
	return New(Sample{
		Messages: map[string]string{
			En: msg,
		},
		Desc:     desc,
		Hint:     hint,
		HttpCode: code,
		Level:    levels.LevelUser,
		Cause:    err,
	})
}
