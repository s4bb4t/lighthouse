package sp

import (
	"errors"
)

// Is returns true if the error is the same as err. Implements the errors.Is
// interface for error comparison. It returns true if the error's cause matches
// the provided error.
func (e *Error) Is(err error) bool {
	if sperr, ok := err.(*Error); ok {
		return sperr.id == e.id
	}
	return errors.Is(e.Core.Cause, err)
}

// DeepIs traverses the entire Error chain to find a matching error. Returns true if
// the provided error is found anywhere in the chain.
func (e *Error) DeepIs(err error) bool {
	var cp = &Error{}
	*cp = *e
	var head *Error
	for {
		head = cp.pop()
		if head == nil {
			break
		}
		if errors.Is(head.Core.Cause, err) {
			return true
		}
	}
	return false
}
