package registry

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"hash"
	"sync"
)

// TODO: arc and create registry system + tests

type (
	// registry stores and manages Error instances with thread-safe access.
	// it uses a hash-based mapping to store errors and includes a mutex for concurrent operations.
	registry struct {
		errs map[int]*sp.Error
		sync.RWMutex
	}
)

// Reg registers an error in the registry.
// It returns a hash id of the error.
// If the error is already registered, it returns the hash id of the existing error.
// If the error is nil, it returns an error.
// The error is validated and stored in the registry.
func (r *registry) Reg(e error) (hash.Hash, error) {
	r.Lock()
	defer r.Unlock()

	if e == nil {
		return nil, sp.New(sp.Sample{
			Messages: map[string]string{
				sp.En: "Nil error provided",
				sp.Ru: "передана nil ошибка",
			},
			Desc:     "Provided error is nil. This is not allowed)",
			Hint:     "Please, check your code and provide a valid error",
			HttpCode: 400,
			Level:    levels.LevelError,
		})
	}

	//h, err := e.done()
	//if err != nil {
	//	return nil, sp.New(sp.Sample{
	//		Messages: map[string]string{
	//			sp.En: "Failed to validate Error",
	//			sp.Ru: "Ошибка в процессе валидации",
	//		},
	//		Desc:     "Failed to create hash id of your error. It happens when you try to register an error with an empty description. Provided data of error in Meta",
	//		Hint:     "Please, check your fields and provide a valid description, hint and EN message for your error",
	//		HttpCode: 400,
	//		Level:    levels.LevelError,
	//		Cause:    err,
	//	})
	//}
	//
	//r.errs[h] = e
	return nil, nil
}

// Get returns an error by its hash id.
// If the error is not found, it returns nil.
// The returned error is a copy of the original error.
func (r *registry) Get(id int) error {
	r.RLock()
	defer r.RUnlock()

	e, ok := r.errs[id]
	if !ok {
		return nil
	}

	cp := &sp.Error{}
	*cp = *sp.Ensure(e)

	return cp.Path(1)
}
