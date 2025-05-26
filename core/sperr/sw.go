package sperr

import "fmt"

// Wrap wraps src into e's underlying SPError
func (e *SPError) Wrap(src *SPError) *SPError {
	if e.IsSP(src) {
		return e
	}

	e.underlying = src
	e.remainsUnderlying = src.remainsUnderlying + 1
	return e
}

// Wrap wraps err into new-initialized SPError from provided Fields
func Wrap(err *SPError, f Fields) *SPError {
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

func (e *SPError) Spin(lvl ErrorLevel) *SPError {
	var cp = &SPError{}
	*cp = *e

	var head, ls *SPError
	head = cp.Pop()
	ls = head
	if head.level > lvl {
		return nil
	}

	cnt := 0
	switch lvl {
	case LevelNoop:
		return nil
	case LevelHighUser, LevelMediumUser, LevelLowUser:
		// TODO: make hints for users only
		// TODO: user data retrieving from meta
		return ls
	case LevelInfo, LevelWarn, LevelError:
		return ls
	case LevelHighDebug, LevelMediumDebug, LevelDeepDebug:
		// TODO: stack trace
		// TODO: All provided data from meta
		for head.level <= lvl {
			ls = head

			// TODO: Log ---------------------------------------------------
			for range cnt {
				fmt.Print(" ")
			}
			fmt.Printf("%s: %s: %s", head.path, head.desc, head.hint)
			fmt.Println()
			// TODO: Log ---------------------------------------------------

			head = cp.Pop()
			if head == nil {
				break
			}
			cnt += 5
		}
		return ls
	default:
		return ls
	}
}
