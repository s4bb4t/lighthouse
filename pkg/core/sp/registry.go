package sp

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"hash"
	"sync"
)

type (
	// registry stores and manages Error instances with thread-safe access.
	// it uses a hash-based mapping to store errors and includes a mutex for concurrent operations.
	registry struct {
		errs map[hash.Hash]*Error
		sync.RWMutex
	}
)

// Registry is a global registry of all errors.
// It is used to store and retrieve errors by their hash id.
// It is thread-safe and can be used concurrently.
// You can add your own errors to the registry using the Reg() method.
// You can retrieve an error by its hash id using the Get() method.
// The registry is initialized with a set of default errors: check Internal, NotFound, BadRequest, Unauthorized, Forbidden, Timeout.
var Registry registry

var (
	// Internal is a hash id of the Internal server error.
	Internal hash.Hash
	// NotFound is a hash id of the Not found error.
	NotFound hash.Hash
	// BadRequest is a hash id of the Bad request error.
	BadRequest hash.Hash
	// Unauthorized is a hash id of the Unauthorized error.
	Unauthorized hash.Hash
	// Forbidden is a hash id of the Forbidden error.
	Forbidden hash.Hash
	// Timeout is a hash id of the Request timeout error.
	Timeout hash.Hash
)

func init() {
	Registry.errs = make(map[hash.Hash]*Error)

	Internal, _ = Registry.Reg(New(Sample{
		Messages: map[string]string{
			En: "Internal server error",
			Ru: "Ошибка сервера",
		},
		Desc:     "Internal server error. We are sorry for the inconvenience",
		Hint:     "Please try again later - we are working on it",
		HttpCode: 500,
		Level:    levels.LevelUser,
	}))

	NotFound, _ = Registry.Reg(New(Sample{
		Messages: map[string]string{
			En: "Resource not found",
			Ru: "Ресурс не найден",
		},
		Desc:     "The requested resource could not be found on this server",
		Hint:     "Please check the URL and try again",
		HttpCode: 404,
		Level:    levels.LevelUser,
	}))

	BadRequest, _ = Registry.Reg(New(Sample{
		Messages: map[string]string{
			En: "Bad request",
			Ru: "Неверный запрос",
		},
		Desc:     "The request could not be understood by the server due to malformed syntax",
		Hint:     "Please check your request parameters and try again",
		HttpCode: 400,
		Level:    levels.LevelUser,
	}))

	Unauthorized, _ = Registry.Reg(New(Sample{
		Messages: map[string]string{
			En: "Unauthorized",
			Ru: "Не авторизован",
		},
		Desc:     "Authentication is required and has failed or has not been provided",
		Hint:     "Please provide valid authentication credentials",
		HttpCode: 401,
		Level:    levels.LevelUser,
	}))

	Forbidden, _ = Registry.Reg(New(Sample{
		Messages: map[string]string{
			En: "Forbidden",
			Ru: "Доступ запрещен",
		},
		Desc:     "You don't have permission to access this resource",
		Hint:     "Please contact your administrator if you need access",
		HttpCode: 403,
		Level:    levels.LevelUser,
	}))

	Timeout, _ = Registry.Reg(New(Sample{
		Messages: map[string]string{
			En: "Request timeout",
			Ru: "Время ожидания истекло",
		},
		Desc:     "The server timed out waiting for the request",
		Hint:     "Please try again. If the problem persists, contact support",
		HttpCode: 408,
		Level:    levels.LevelUser,
	}))
}

// Reg registers an error in the registry.
// It returns a hash id of the error.
// If the error is already registered, it returns the hash id of the existing error.
// If the error is nil, it returns an error.
// The error is validated and stored in the registry.
func (r *registry) Reg(e *Error) (hash.Hash, error) {
	r.Lock()
	defer r.Unlock()

	if e == nil {
		return nil, New(Sample{
			Messages: map[string]string{
				En: "Nil error provided",
				Ru: "передана nil ошибка",
			},
			Desc:     "Provided error is nil. This is not allowed)",
			Hint:     "Please, check your code and provide a valid error",
			HttpCode: 400,
			Level:    levels.LevelError,
		})
	}

	h, err := e.done()
	if err != nil {
		return nil, New(Sample{
			Messages: map[string]string{
				En: "Failed to validate Error",
				Ru: "Ошибка в процессе валидации",
			},
			Desc:     "Failed to create hash id of your error. It happens when you try to register an error with an empty description. Provided data of error in Meta",
			Hint:     "Please, check your fields and provide a valid description, hint and EN message for your error",
			HttpCode: 400,
			Level:    levels.LevelError,
			Cause:    err,
			Meta: map[string]any{
				SPErrorKey: *e,
			},
		})
	}

	r.errs[h] = e
	return h, nil
}

// Get returns an error by its hash id.
// If the error is not found, it returns nil.
// The returned error is a copy of the original error.
func (r *registry) Get(h hash.Hash) *Error {
	r.RLock()
	defer r.RUnlock()

	sp, ok := r.errs[h]
	if !ok {
		return nil
	}

	cp := &Error{}
	*cp = *sp

	return cp._path(1)
}
