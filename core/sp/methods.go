package sp

import (
	"crypto/sha256"
	"fmt"
	"github.com/s4bb4t/lighthouse/core/levels"
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

// Done finalizes the error handling process and returns a hash and an error if any occurs during the operation.
// It generates a hash ID based on the Error's fields
// Error can't be used without calling Done() or MustDone()
func (e *Error) Done() (hash.Hash, error) {
	_ = e._path(1)
	return e.done()
}

// MustDone generates a hash ID based on the Error's fields
// Error can't be used without calling Done() or MustDone()
// It panics if an error is encountered during the process.
func (e *Error) MustDone() *Error {
	if _, err := e.done(); err != nil {
		panic(err)
	}
	return e._path(1)
}

// Error returns the Error's description.
func (e *Error) Error() string {
	return e.desc + ": " + e.hint
}

func (e *Error) Unwrap() error {
	if e.underlying != nil {
		return e.underlying
	}

	return e.cause
}

// Cast attempts to convert a generic error to *Error type.
func Cast(err error) (*Error, bool) {
	e, b := err.(*Error)
	return e, b
}

// AllMeta returns a copy of all metadata associated with the error.
// The returned map is a new instance to prevent modification of the original metadata.
func (e *Error) AllMeta() map[string]any {
	meta := make(map[string]any)
	maps.Copy(meta, e.meta)
	return meta
}

// ReadCaused returns the underlying cause of the error.
// If there is no cause, it returns nil.
func (e *Error) ReadCaused() error {
	return e.cause
}

// ReadMsg returns the error message for the specified language code.
// Parameter lg represents the language code to retrieve the message for.
func (e *Error) ReadMsg(lg string) string {
	return e.messages[lg]
}

// ReadDesc returns the description of the error.
// The description provides additional context about the error.
func (e *Error) ReadDesc() string {
	return e.desc
}

// ReadHint returns a hint or suggestion related to resolving the error.
// The hint provides guidance on how to fix or handle the error.
func (e *Error) ReadHint() string {
	return e.hint
}

// ReadCode returns the HTTP status code associated with the error.
// This code indicates the type of error in HTTP.
func (e *Error) ReadCode() int {
	return e.httpCode
}

// ReadLevel returns the severity level of the error.
// The level indicates how critical or severe the error is.
func (e *Error) ReadLevel() levels.Level {
	return e.level
}

// ReadMeta returns the metadata associated with the error.
// The metadata provides additional information about the error.
func (e *Error) ReadMeta(key string) any {
	return e.meta[key]
}
func (e *Error) ReadPath() string {
	return e.path
}

// ReadHash returns the hash ID of the error.
// The hash ID is a unique identifier for the error.
func (e *Error) ReadHash() hash.Hash {
	return e.id
}
func (e *Error) ReadTime() time.Time {
	return e.timestamp
}
