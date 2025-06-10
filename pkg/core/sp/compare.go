package sp

import (
	"errors"
	"strings"
)

// Is checks if the provided error is a member of the Error chain.
func (e *Error) Is(err error) bool {
	switch v := err.(type) {
	case *Error:
		if strings.ToLower(e.Msg(En)) == strings.ToLower(v.Msg(En)) || v.Desc() == e.Desc() {
			return true
		}
	default:
		return errors.Is(e.Core.Cause, err)
	}
	return false
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
