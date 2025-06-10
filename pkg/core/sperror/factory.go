// Package sperror provides structured error handling functionality with support for
// localized messages, error chaining, and detailed error information tracking.
package sperror

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"maps"
	"path/filepath"
	"runtime"
)

type (
	CoreError struct {
		Desc   string // detailed description
		Hint   string // how to resolve
		Source string // source of error (file path, line number, etc.)
		Cause  error  // nested error
	}

	UserError struct {
		Messages map[string]string // localized message
		HttpCode int               // HTTP status
		Level    levels.Level      // error level
	}

	// Error represents a structured error type with extended information and metadata.
	// It supports localized messages, detailed descriptions, resolution hints, and additional context.
	Error struct {
		Core CoreError
		User UserError

		meta map[string]any // arbitrary fields (user_id, trace_id, etc.)

		remainsUnderlying int
		underlying        *Error
	}
)

// NewSpErr creates and returns a new instance of Error.
// It initializes an empty Error struct that can be further configured using method chaining.
// This is the base constructor for creating new structured errors.
func NewSpErr() *Error {
	return &Error{
		User: UserError{
			Messages: make(map[string]string),
		},
		meta: make(map[string]any),
	}
}

// New constructs and returns a new Error based on the provided Sample.
// It initializes a new Error instance, then copies all fields from the provided Sample and returns
// the configured and ready to use Error instance
func New(s Sample) *Error {
	sp := NewSpErr()

	if sp.User.Messages == nil {
		sp.User.Messages = make(map[string]string)
	}
	maps.Copy(sp.User.Messages, s.Messages)

	if sp.meta == nil {
		sp.meta = make(map[string]any)
	}
	maps.Copy(sp.meta, s.Meta)

	return sp.
		SetDesc(s.Desc).
		SetHint(s.Hint).
		SetCode(s.HttpCode).
		SetLevel(s.Level).
		SetCaused(s.Cause).
		path(1)
}

// SetCaused sets the underlying error.
func (e *Error) SetCaused(err error) *Error {
	e.Core.Cause = err
	return e
}

// SetMsg sets the localized message for the given language.
func (e *Error) SetMsg(lg, msg string) *Error {
	if e.User.Messages == nil {
		e.User.Messages = make(map[string]string)
	}

	e.User.Messages[lg] = msg
	return e
}

// SetDesc sets the complete description for the given language.
func (e *Error) SetDesc(desc string) *Error {
	e.Core.Desc = desc
	return e
}

// SetHint sets the hint for the given language.
func (e *Error) SetHint(hint string) *Error {
	e.Core.Hint = hint
	return e
}

// SetCode sets the HTTP status code for the error.
// It accepts an integer representing the HTTP status code and returns the modified Error.
func (e *Error) SetCode(httpCode int) *Error {
	e.User.HttpCode = httpCode
	return e
}

// SetLevel sets the severity level of the error.
// It accepts a Level value and returns the modified Error.
func (e *Error) SetLevel(lvl levels.Level) *Error {
	e.User.Level = lvl
	return e
}

// AddMeta adds a key-value pair to the error's metadata.
// It accepts a string key and any value, returning the modified Error.
func (e *Error) AddMeta(key string, val any) *Error {
	e.meta[key] = val
	return e
}

// path is an internal method that sets the error source based on the caller's location
// Stack frame level to look up (relative to caller)
// Returns:
// - *Error: The modified error instance with source set
// The source format is "absolute_file_path:line_number"
func (e *Error) path(lvl int) *Error {
	_, file, line, ok := runtime.Caller(lvl + 1)
	if ok {
		absPath, err := filepath.Abs(file)
		if err != nil {
			panic(err)
		}
		e.Core.Source = fmt.Sprintf("%s:%d", absPath, line)
	}
	return e
}

// SetSource sets the error source based on the caller's location
// Stack frame level to look up (relative to caller)
// Returns:
// - *Error: The modified error instance with source set
// The source format is "absolute_file_path:line_number"
func (e *Error) SetSource() *Error {
	return e.path(1)
}

// HelperSetSource sets the error source based on the caller's location + 1
// It needs only if you want to set correct source of the error that is created not in place of usage
func (e *Error) HelperSetSource() *Error {
	return e.path(2)
}
