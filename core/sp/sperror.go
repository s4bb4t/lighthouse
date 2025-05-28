package sp

import (
	"crypto/sha256"
	"fmt"
	"github.com/s4bb4t/lighthouse/core/levels"
	"hash"
	"maps"
	"time"
)

type (
	// SPError represents a structured error type with extended information and metadata.
	// It supports localized messages, detailed descriptions, resolution hints, and additional context.
	SPError struct {
		messages map[string]string // localized message
		desc     string            // detailed description
		hint     string            // how to resolve
		path     string            // path/operation

		id        hash.Hash    // UUID or content hash
		httpCode  int          // HTTP status
		level     levels.Level // error level
		timestamp time.Time    // when occurred

		cause error          // nested error
		stack []string       // stack trace
		meta  map[string]any // arbitrary fields (user_id, trace_id, etc.)

		remainsUnderlying int
		underlying        *SPError
	}
)

// NewSpErr creates and returns a new instance of SPError.
// It initializes an empty SPError struct that can be further configured using method chaining.
func NewSpErr() *SPError {
	return &SPError{}
}

// SP constructs and returns a new SPError based on the provided Err, copying messages and meta along with validation.
func SP(f Err) *SPError {
	sp := NewSpErr()

	if sp.messages == nil {
		sp.messages = make(map[string]string)
	}

	maps.Copy(sp.messages, f.Messages)
	if sp.meta == nil {
		sp.meta = make(map[string]any)
	}
	maps.Copy(sp.meta, f.Meta)

	_, err := sp.
		Path(f.Path).
		Desc(f.Desc).
		Hint(f.Hint).
		Code(f.HttpCode).
		Level(f.Level).
		Caused(f.Cause).
		Done()
	if err != nil {
		panic(err)
	}
	return sp
}

// Path appends the operation name to the path.
func (e *SPError) Path(path string) *SPError {
	e.path = path + e.path
	return e
}

// Caused sets the underlying error.
func (e *SPError) Caused(err error) *SPError {
	e.cause = err
	return e
}

// Msg sets the localized message for the given language.
func (e *SPError) Msg(lg, msg string) *SPError {
	e.messages[lg] = msg
	return e
}

// Desc sets the complete description for the given language.
func (e *SPError) Desc(desc string) *SPError {
	e.desc = desc
	return e
}

// Hint sets the hint for the given language.
func (e *SPError) Hint(hint string) *SPError {
	e.hint = hint
	return e
}

// Code sets the HTTP status code for the error.
// It accepts an integer representing the HTTP status code and returns the modified SPError.
func (e *SPError) Code(httpCode int) *SPError {
	e.httpCode = httpCode
	return e
}

// Level sets the severity level of the error.
// It accepts an Level value and returns the modified SPError.
func (e *SPError) Level(lvl levels.Level) *SPError {
	e.level = lvl
	return e
}

// Meta adds a key-value pair to the error's metadata.
// It accepts a string key and any value, returning the modified SPError.
func (e *SPError) Meta(key string, val any) *SPError {
	e.meta[key] = val
	return e
}

// Done generates a hash ID based on the SPError's fields and returns any error encountered during the process.
// SPError can't be used without calling Done()
func (e *SPError) Done() (hash.Hash, error) {
	if e == nil || e.desc == "" || e.messages[En] == "" {
		return nil, fmt.Errorf("do not use empty sperror: it may cause misundertstanings")
	}

	e.timestamp = time.Now()
	e.id = sha256.New()
	_, err := e.id.Write([]byte(e.desc + e.hint + e.messages[En]))
	if err != nil {
		return nil, err
	}

	return e.id, err
}

// MustDone generates a hash ID based on the SPError's fields
// SPError can't be used without calling Done()
func (e *SPError) MustDone() hash.Hash {
	if e == nil || e.desc == "" || e.messages[En] == "" {
		panic("do not use empty sperror: it may cause misundertstanings")
	}

	e.timestamp = time.Now()
	e.id = sha256.New()
	_, err := e.id.Write([]byte(e.desc + e.hint + e.messages[En]))
	if err != nil {
		panic(err)
	}
	return e.id
}

// Error returns the SPError's description.
func (e *SPError) Error() string {
	return e.desc + ": " + e.hint
}

func (e *SPError) Unwrap() error {
	return e.cause
}

func Cast(err error) (*SPError, bool) {
	e, b := err.(*SPError)
	return e, b
}

// AllMeta returns a copy of all metadata associated with the error.
// The returned map is a new instance to prevent modification of the original metadata.
func (e *SPError) AllMeta() map[string]any {
	meta := make(map[string]any)
	maps.Copy(meta, e.meta)
	return meta
}

// ReadCaused returns the underlying cause of the error.
// If there is no cause, it returns nil.
func (e *SPError) ReadCaused() error {
	return e.cause
}

// ReadMsg returns the error message for the specified language code.
// Parameter lg represents the language code to retrieve the message for.
func (e *SPError) ReadMsg(lg string) string {
	return e.messages[lg]
}

// ReadDesc returns the description of the error.
// The description provides additional context about the error.
func (e *SPError) ReadDesc() string {
	return e.desc
}

// ReadHint returns a hint or suggestion related to resolving the error.
// The hint provides guidance on how to fix or handle the error.
func (e *SPError) ReadHint() string {
	return e.hint
}

// ReadCode returns the HTTP status code associated with the error.
// This code indicates the type of error in HTTP context.
func (e *SPError) ReadCode() int {
	return e.httpCode
}

// ReadLevel returns the severity level of the error.
// The level indicates how critical or severe the error is.
func (e *SPError) ReadLevel() levels.Level {
	return e.level
}

func (e *SPError) ReadPath() string {
	return e.path
}

func (e *SPError) ReadTime() time.Time {
	return e.timestamp
}
