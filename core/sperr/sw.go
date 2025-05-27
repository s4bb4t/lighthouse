package sperr

import "github.com/s4bb4t/lighthouse/core/levels"

// Wrap wraps src into e's underlying SPError
func (e *SPError) Wrap(src *SPError) *SPError {
	if e.IsSP(src) {
		return e
	}

	e.underlying = src
	e.remainsUnderlying = src.remainsUnderlying + 1
	return e
}

// Wrap wraps err into new-initialized SPError from provided Err or existing SPError.
// It returns dest if src matches its hash, otherwise wraps src into dest.
func Wrap(src *SPError, dst any) *SPError {
	switch v := dst.(type) {
	case Err:
		h, _ := v.hash()
		if cmpHashes(src.id, h) {
			return src
		}

		res := SP(v)
		res.underlying = src
		res.remainsUnderlying = src.remainsUnderlying + 1
		return res
	case *SPError:
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
func (e *SPError) Pop() *SPError {
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

// Spin processes the error chain until reaching a specified error level.
// Returns the last error that matches the specified level.
// If the initial error level is higher than required, returns an Internal error.
func (e *SPError) Spin(lvl levels.ErrorLevel) *SPError {
	if lvl == levels.LevelNoop {
		return nil
	}

	cp := &SPError{}
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
