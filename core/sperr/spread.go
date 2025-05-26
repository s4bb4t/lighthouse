package sperr

import "maps"

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
func (e *SPError) ReadLevel() ErrorLevel {
	return e.level
}
