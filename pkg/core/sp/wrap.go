package sp

// Unwrap returns the underlying error if it exists; otherwise, it returns the cause of the error.
func (e *Error) Unwrap() error {
	if e.underlying != nil {
		return e.underlying
	}
	return e.Core.Cause
}

// Wrap wraps src into e's underlying Error
func (e *Error) Wrap(src *Error) *Error {
	e.underlying = src
	e.remainsUnderlying = src.remainsUnderlying + 1
	return e
}

// Wrap wraps `src` into new-initialized Error from provided Sample or existing Error.
func Wrap(src *Error, dst any) *Error {
	switch v := dst.(type) {
	case Sample:
		res := New(v)
		res.underlying = src
		res.remainsUnderlying = src.remainsUnderlying + 1
		return res.Path(1)
	case *Error:
		v.underlying = src
		v.remainsUnderlying = src.remainsUnderlying + 1
		return v
	default:
		panic("unsupported destination type")
	}
}
