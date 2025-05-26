package sperr

import (
	"crypto/sha256"
	"hash"
)

type Fields struct {
	Messages map[string]string // localized message
	Desc     string            // detailed description
	Hint     string            // how to resolve
	Path     string            // path/operation

	HttpCode int        // HTTP status
	Level    ErrorLevel // error level

	Cause error          // nested error
	Meta  map[string]any // arbitrary fields (user_id, trace_id, etc.)
}

func (f Fields) hash() (hash.Hash, error) {
	h := sha256.New()
	_, err := h.Write([]byte(f.Desc + f.Hint + f.Messages[En]))
	return h, err
}
