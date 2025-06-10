package sperror

import (
	"errors"
)

// Unwrap returns the underlying error if it exists; otherwise, it returns the cause of the error.
func (e *Error) Unwrap() error {
	if e.underlying != nil {
		return e.underlying
	}
	if e.Core.Cause != nil {
		return e.Core.Cause
	}
	return errors.New(e.Error())
}

// Wrap wraps src into e's cause Error
func (e *Error) Wrap(err error) *Error {
	e.Core.Cause = err
	return e
}

// Wrap wraps `src` into existing Error.
func Wrap(src *Error, dst *Error) *Error {
	dst.underlying = src
	dst.remainsUnderlying = src.remainsUnderlying + 1
	return dst
}

// WrapNew wraps `src` into new-initialized Error from provided Sample.
func WrapNew(src *Error, dst Sample) *Error {
	res := New(dst)
	res.underlying = src
	res.remainsUnderlying = src.remainsUnderlying + 1
	return res.path(1)
}
