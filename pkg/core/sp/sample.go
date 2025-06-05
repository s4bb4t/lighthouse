package sp

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
)

type Sample struct {
	Messages map[string]string // localized message
	Desc     string            // detailed description
	Hint     string            // how to resolve

	HttpCode int          // HTTP status
	Level    levels.Level // error level

	Cause error          // nested error
	Meta  map[string]any // arbitrary fields (user_id, trace_id, etc.)
}
