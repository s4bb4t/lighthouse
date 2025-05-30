// Package sp provides structured error handling functionality with support for
// localized messages, error chaining, and detailed error information tracking.
package sp

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/core/levels"
	"hash"
	"maps"
	"path/filepath"
	"runtime"
	"time"
)

type (
	// Error represents a structured error type with extended information and metadata.
	// It supports localized messages, detailed descriptions, resolution hints, and additional context.
	Error struct {
		messages map[string]string // localized message
		desc     string            // detailed description
		hint     string            // how to resolve
		source   string            // source of error (file path, line number, etc.)

		id        hash.Hash    // UUID or content hash
		httpCode  int          // HTTP status
		level     levels.Level // error level
		timestamp time.Time    // when occurred

		cause error          // nested error
		meta  map[string]any // arbitrary fields (user_id, trace_id, etc.)

		remainsUnderlying int
		underlying        *Error
	}
)

// NewSpErr creates and returns a new instance of Error.
// It initializes an empty Error struct that can be further configured using method chaining.
// This is the base constructor for creating new structured errors.
func NewSpErr() *Error {
	return &Error{}
}

// New constructs and returns a new Error based on the provided Err.
// It initializes a new Error instance, then copies all fields from the provided Err and returns
// the configured and ready to use Error instance
func New(f Err) *Error {
	sp := NewSpErr()

	if sp.messages == nil {
		sp.messages = make(map[string]string)
	}
	maps.Copy(sp.messages, f.Messages)

	if sp.meta == nil {
		sp.meta = make(map[string]any)
	}
	maps.Copy(sp.meta, f.Meta)

	return sp.
		_path(1).
		Desc(f.Desc).
		Hint(f.Hint).
		Code(f.HttpCode).
		Level(f.Level).
		Caused(f.Cause).
		mustDone()
}

// Caused sets the underlying error.
func (e *Error) Caused(err error) *Error {
	e.cause = err
	return e
}

// Msg sets the localized message for the given language.
func (e *Error) Msg(lg, msg string) *Error {
	if e.messages == nil {
		e.messages = make(map[string]string)
	}

	e.messages[lg] = msg
	return e
}

// Desc sets the complete description for the given language.
func (e *Error) Desc(desc string) *Error {
	e.desc = desc
	return e
}

// Hint sets the hint for the given language.
func (e *Error) Hint(hint string) *Error {
	e.hint = hint
	return e
}

// Code sets the HTTP status code for the error.
// It accepts an integer representing the HTTP status code and returns the modified Error.
func (e *Error) Code(httpCode int) *Error {
	e.httpCode = httpCode
	return e
}

// Level sets the severity level of the error.
// It accepts a Level value and returns the modified Error.
func (e *Error) Level(lvl levels.Level) *Error {
	e.level = lvl
	return e
}

// Meta adds a key-value pair to the error's metadata.
// It accepts a string key and any value, returning the modified Error.
func (e *Error) Meta(key string, val any) *Error {
	e.meta[key] = val
	return e
}

// _path is an internal method that sets the error source based on the caller's location
// Stack frame level to look up (relative to caller)
// Returns:
// - *Error: The modified error instance with source set
// The source format is "absolute_file_path:line_number"
func (e *Error) _path(lvl int) *Error {
	_, file, line, ok := runtime.Caller(lvl + 1)
	if ok {
		absPath, err := filepath.Abs(file)
		if err != nil {
			panic(err)
		}
		e.source = fmt.Sprintf("%s:%d", absPath, line)
	}
	return e
}

// Source sets the error source based on the caller's location
// Stack frame level to look up (relative to caller)
// Returns:
// - *Error: The modified error instance with source set
// The source format is "absolute_file_path:line_number"
func (e *Error) Source() *Error {
	return e._path(0)
}
