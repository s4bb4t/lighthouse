package sp

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
)

// Pop extracts and returns the current error from the error chain.
// If the current object is nil or there are no more underlying errors (remainsUnderlying == -1),
// returns nil.
func (e *Error) pop() *Error {
	if e == nil || e.remainsUnderlying == -1 {
		return nil
	}
	result := *e
	result.underlying = nil
	if e.underlying != nil {
		*e = *e.underlying
	} else {
		e.remainsUnderlying--
	}
	return &result
}

// Spin returns the most relevant *Error instance from the error chain that matches the provided severity level.
//
// This method traverses the error chain (created via Wrap) and returns the last *Error whose Level is
// less than or equal to the provided level. This allows you to present different error messages to users,
// logs, or monitoring systems depending on the trust/context level.
//
// For example:
//   - LevelUser → return user-friendly message
//   - LevelDebug → return detailed system context
//   - LevelError → return technical issue without leaking internals
//
// ⚠️ If Spin() is not called and the error is passed as-is, the most recent (outer) error will be returned by default.
//
// Usage pattern:
//
//	if err := handler.Do(); err != nil {
//	    return sp.Ensure(err).Spin(levels.LevelUser)
//	}
//
// If no error matches the level, Spin returns a generic internal error.
//
// Recommended practices:
// - Always wrap each layer’s error with Wrap()
// - Use Spin(level) before rendering/logging to avoid leaking internals
// - Don’t call Spin() in places where full technical info is needed (e.g., logs)
func (e *Error) Spin(lvl levels.Level) *Error {
	if lvl == levels.LevelNoop {
		return nil
	}

	cp := &Error{}
	*cp = *e
	cur := cp.pop()
	if cur == nil {
		return nil
	}

	if cur.User.Level > lvl {
		return New(Sample{
			Messages: map[string]string{
				En: "No error found for level",
			},
			Desc: "Level provided is higher than the level of the error",
			Hint: "Please, check your code and provide a valid error level",
		}).path(1)
	}

	last := cur
	for cur != nil && cur.User.Level <= lvl {
		last, cur = cur, cp.pop()
	}
	return last
}
