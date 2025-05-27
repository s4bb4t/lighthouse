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

// Wrap wraps err into new-initialized SPError from provided Err
func Wrap(err *SPError, f Err) *SPError { // TODO: generic
	h, _ := f.hash()
	if cmpHashes(err.id, h) {
		return err
	}

	sp := SP(f)
	sp.underlying = err
	sp.remainsUnderlying = err.remainsUnderlying + 1
	return sp
}

func (e *SPError) Pop() *SPError {
	if e == nil || e.remainsUnderlying == -1 {
		return nil
	}

	r := *e
	r.underlying = nil

	if e.underlying != nil {
		*e = *e.underlying
	} else {
		e.remainsUnderlying--
	}
	return &r
}

func (e *SPError) Spin(lvl levels.ErrorLevel) *SPError {
	var cp = &SPError{}
	*cp = *e

	var head, ls *SPError
	head = cp.Pop()
	ls = head
	if head.level > lvl {
		return Registry.errs[Internal]
	}

	cnt := 0
	switch lvl {
	case levels.LevelNoop:
		return nil
	default:
		for head.level <= lvl {
			ls = head
			head = cp.Pop()
			if head == nil {
				break
			}
			cnt += 5
		}
		return ls
	}
}
