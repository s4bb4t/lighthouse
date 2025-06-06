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

// Spin returns a copy
// of the last available error for the selected level. If you pass Error somewhere without calling
// Spin() - the last error in the chain will be passed
//
// Spin "unwinds" the Error up to the specified error level. With this function you can interpret the
// same error differently depending on the context - lvl (error level).
//
// If no suitable level is found in this Error instance - Registry[Internal] will be returned
//
// Recommended:
//
// - wrap the error through Wrap at each layer of your application to avoid losing context or
// transmitting confidential information that may be contained within Error.
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
		})
	}

	last := cur
	for cur != nil && cur.User.Level <= lvl {
		last, cur = cur, cp.pop()
	}
	return last
}
