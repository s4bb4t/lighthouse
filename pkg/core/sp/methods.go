package sp

import (
	"crypto/sha256"
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"hash"
	"maps"
	"time"
)

// done performs finalization of the Error instance by:
// Validating required fields (description and English message)
// Setting timestamp
// Generating unique hash ID based on error content
// Returns the generated hash or error if validation fails
// Internal method used by Done() and MustDone()
func (e *Error) done() (hash.Hash, error) {
	if e == nil || e.desc == "" || e.messages[En] == "" {
		return nil, fmt.Errorf("do not use empty sperror: it may cause misunderstandings")
	}

	e.timestamp = time.Now()
	e.id = sha256.New()
	_, err := e.id.Write([]byte(e.desc + e.hint + e.messages[En]))
	if err != nil {
		return nil, err
	}
	return e.id, err
}

// mustDone ensures the completion of an operation, panicking if an error occurs, and returns the updated Error instance.
func (e *Error) mustDone() *Error {
	if _, err := e.done(); err != nil {
		panic(err)
	}
	return e._path(1)
}

// Error returns the Error's description.
func (e *Error) Error() string {
	return e.desc + ": " + e.hint
}

// cast attempts to convert a generic error to *Error type.
func cast(err error) (*Error, bool) {
	e, b := err.(*Error)
	return e, b
}

// Ensure wraps a given error into a custom *Error type if it is not already of that type.
// If the error is already of the *Error type, it is returned as-is.
// This method is useful for ensuring that an error is of the *Error type.
// It is recommended to use this method instead of casting the error to *Error type directly.
func Ensure(err error) *Error {
	if serr, ok := cast(err); ok {
		return serr
	}
	return New(Sample{
		Messages: map[string]string{En: "Unknown error"},
		Desc:     err.Error(),
		Hint:     "Check original .Error()",
		Level:    levels.LevelError,
		Cause:    err,
	})._path(1)
}

// AllMeta returns a copy of all metadata associated with the error.
// The returned map is a new instance to prevent modification of the original metadata.
func (e *Error) AllMeta() map[string]any {
	meta := make(map[string]any)
	maps.Copy(meta, e.meta)
	return meta
}

// Caused returns the underlying cause of the error.
// If there is no cause, it returns nil.
func (e *Error) Caused() error {
	return e.cause
}

// Msg returns the error message for the specified language code.
// Parameter lg represents the language code to retrieve the message for.
func (e *Error) Msg(lg string) string {
	return e.messages[lg]
}

// Desc returns the description of the error.
// The description provides additional context about the error.
func (e *Error) Desc() string {
	return e.desc
}

// Hint returns a hint or suggestion related to resolving the error.
// The hint provides guidance on how to fix or handle the error.
func (e *Error) Hint() string {
	return e.hint
}

// Code returns the HTTP status code associated with the error.
// This code indicates the type of error in HTTP.
func (e *Error) Code() int {
	return e.httpCode
}

// Level returns the severity level of the error.
// The level indicates how critical or severe the error is.
func (e *Error) Level() levels.Level {
	return e.level
}

// Meta returns the metadata associated with the error.
// The metadata provides additional information about the error.
func (e *Error) Meta(key string) any {
	return e.meta[key]
}

// Source retrieves the source field value from the Error instance. It returns the source as a string.
func (e *Error) Source() string {
	return e.source
}

// Hash returns the hash ID of the error.
// The hash ID is a unique identifier for the error.
func (e *Error) Hash() hash.Hash {
	return e.id
}

// Time returns the timestamp associated with the Error instance.
func (e *Error) Time() time.Time {
	return e.timestamp
}

// Stack returns the stack trace of the error.
// Stack trace can be accessed ONLY after Spin() is used
func (e *Error) Stack() []string {
	return e.stackTrace
}
