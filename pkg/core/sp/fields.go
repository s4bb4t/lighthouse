package sp

import (
	"crypto/sha256"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"hash"
)

type Err struct {
	Messages map[string]string // localized message
	Desc     string            // detailed description
	Hint     string            // how to resolve

	HttpCode int          // HTTP status
	Level    levels.Level // error level

	Cause error          // nested error
	Meta  map[string]any // arbitrary fields (user_id, trace_id, etc.)
}

func (f Err) hash() (hash.Hash, error) {
	h := sha256.New()
	_, err := h.Write([]byte(f.Desc + f.Hint + f.Messages[En]))
	return h, err
}
