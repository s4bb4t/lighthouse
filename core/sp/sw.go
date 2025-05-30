package sp

import (
	"github.com/s4bb4t/lighthouse/core/levels"
)

// TODO: implement function that retrieve whole source of error based on Level

// Wrap wraps src into e's underlying Error
func (e *Error) Wrap(src *Error) *Error {
	if e.IsSP(src) {
		return e
	}

	e.underlying = src
	e.remainsUnderlying = src.remainsUnderlying + 1
	return e
}

// Wrap wraps `src` into new-initialized Error from provided Err or existing Error.
// It returns dest if src matches its hash, otherwise wraps src into dest.
func Wrap(src *Error, dst any) *Error {
	switch v := dst.(type) {
	case Err:
		h, _ := v.hash()
		if src.id == nil {
			panic("source error is not validated through Done()")
		}
		if cmpHashes(src.id, h) {
			return src
		}

		res := New(v)
		res.underlying = src
		res.remainsUnderlying = src.remainsUnderlying + 1

		if _, err := res.done(); err != nil {
			panic(err)
		}
		return res
	case *Error:
		if src.IsSP(v) {
			return src
		}

		v.underlying = src
		v.remainsUnderlying = src.remainsUnderlying + 1
		return v
	default:
		panic("unsupported destination type")
	}
}

// Pop extracts and returns the current error from the error chain.
// If the current object is nil or there are no more underlying errors (remainsUnderlying == -1),
// returns nil.
func (e *Error) Pop() *Error {
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

	cur := cp.Pop()
	if cur == nil {
		return nil
	}
	if cur.level > lvl {
		return Registry.errs[Internal]
	}

	last := cur
	for cur != nil && cur.level <= lvl {
		last, cur = cur, cp.Pop()
	}

	return last
}
