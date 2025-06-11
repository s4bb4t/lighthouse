package sperror

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"maps"
)

// Error returns the Error's description.
func (e *Error) Error() string {
	return e.Core.Desc + ": " + e.Core.Hint
}

// Copy returns a copy of the Error instance.
func (e *Error) Copy() *Error {
	var err Error
	err = *e
	return &err
}

// Copy returns a copy of the Error instance.
func Copy(e *Error) *Error {
	return e.Copy()
}

// Ensure wraps a given error into a custom *Error type if it is not already of that type.
// If the error is already of the *Error type, it is returned as-is.
// This method is useful for ensuring that an error is of the *Error type.
// It is recommended to use this method instead of casting the error to an *Error type directly.
func Ensure(err error) *Error {
	if sperr, ok := err.(*Error); ok {
		return sperr
	}
	return New(Sample{
		Messages: map[string]string{En: "Unknown error"},
		Desc:     err.Error(),
		Hint:     "Check original .Error()",
		Level:    levels.LevelError,
		Cause:    err,
	}).path(1)
}

// AllMeta returns a copy of all metadata associated with the error.
// The returned map is a new instance to prevent modification of the original metadata.
func (e *Error) AllMeta() map[string]any {
	meta := make(map[string]any)
	if e.meta == nil {
		return meta
	}
	maps.Copy(meta, e.meta)
	return meta
}

// Caused returns the underlying cause of the error.
// If there is no cause, it returns nil.
func (e *Error) Caused() error {
	return e.Core.Cause
}

// Msg returns the error message for the specified language code.
// Parameter lg represents the language code to retrieve the message for.
func (e *Error) Msg(lg string) string {
	return e.User.Messages[lg]
}

// Desc returns the description of the error.
// The description provides additional context about the error.
func (e *Error) Desc() string {
	return e.Core.Desc
}

// Hint returns a hint or suggestion related to resolving the error.
// The hint provides guidance on how to fix or handle the error.
func (e *Error) Hint() string {
	return e.Core.Hint
}

// Code returns the HTTP status code associated with the error.
// This code indicates the type of error in HTTP.
func (e *Error) Code() int {
	return e.User.HttpCode
}

// Level returns the severity level of the error.
// The level indicates how critical or severe the error is.
func (e *Error) Level() levels.Level {
	return e.User.Level
}

// Meta returns the metadata associated with the error.
// The metadata provides additional information about the error.
func (e *Error) Meta(key string) any {
	if e.meta == nil {
		return nil
	}
	return e.meta[key]
}

// Source retrieves the source field value from the Error instance. It returns the source as a string.
func (e *Error) Source() string {
	return e.Core.Source
}
