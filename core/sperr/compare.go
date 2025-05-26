package sperr

import (
	"errors"
	"hash"
	"slices"
)

// Is returns true if the error is the same as err. Implements the errors.Is
// interface for error comparison. It returns true if the error's cause matches
// the provided error.
func (e *SPError) Is(err error) bool {
	return errors.Is(e.cause, err)
}

// IsSP compares two SPErrors by their hash IDs.
// It returns true if both errors have the same hash ID.
// IsSP compares hashes of SpErrors, not their values or descriptions.
func (e *SPError) IsSP(err *SPError) bool {
	return slices.Compare(e.id.Sum(nil), err.id.Sum(nil)) == 0
}

func cmpHashes(h1, h2 hash.Hash) bool {
	return slices.Compare(h1.Sum(nil), h2.Sum(nil)) == 0
}
