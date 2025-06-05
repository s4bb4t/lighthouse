package sp

// Unwrap returns the underlying error if it exists; otherwise, it returns the cause of the error.
func (e *Error) Unwrap() error {
	if e.underlying != nil {
		return e.underlying
	}
	return e.cause
}

// Wrap wraps src into e's underlying Error
func (e *Error) Wrap(src *Error) *Error {
	if e.IsSP(src) {
		return e
	}

	e.underlying = src
	e.remainsUnderlying = src.remainsUnderlying + 1
	return e
}

// Wrap wraps `src` into new-initialized Error from provided Sample or existing Error.
func Wrap(src *Error, dst any) *Error {
	switch v := dst.(type) {
	case Sample:
		if src == nil || src.id == nil {
			panic("source error is not validated through Done()")
		}

		res := New(v)
		res.underlying = src
		res.remainsUnderlying = src.remainsUnderlying + 1

		if _, err := res.done(); err != nil {
			panic(err)
		}
		return res._path(1)
	case *Error:
		v.underlying = src
		v.remainsUnderlying = src.remainsUnderlying + 1
		return v
	default:
		panic("unsupported destination type")
	}
}
