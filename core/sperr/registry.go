package sperr

import (
	"hash"
	"sync"
)

type (
	// registry stores and manages SPError instances with thread-safe access.
	// it uses a hash-based mapping to store errors and includes a mutex for concurrent operations.
	registry struct {
		errs map[hash.Hash]*SPError
		sync.RWMutex
	}
)

var Registry registry
var (
	Internal     hash.Hash
	NotFound     hash.Hash
	BadRequest   hash.Hash
	Unauthorized hash.Hash
	Forbidden    hash.Hash
	Timeout      hash.Hash
)

func init() {
	Registry.errs = make(map[hash.Hash]*SPError)

	Internal, _ = Registry.Reg(SP(Err{
		Messages: map[string]string{
			En: "Internal server error",
			Ru: "Ошибка сервера",
		},
		Desc:     "Internal server error. We are sorry for the inconvenience",
		Hint:     "Please try again later - we are working on it",
		Path:     "",
		HttpCode: 500,
		Level:    LevelHighUser,
	}))

	NotFound, _ = Registry.Reg(SP(Err{
		Messages: map[string]string{
			En: "Resource not found",
			Ru: "Ресурс не найден",
		},
		Desc:     "The requested resource could not be found on this server",
		Hint:     "Please check the URL and try again",
		Path:     "",
		HttpCode: 404,
		Level:    LevelHighUser,
	}))

	BadRequest, _ = Registry.Reg(SP(Err{
		Messages: map[string]string{
			En: "Bad request",
			Ru: "Неверный запрос",
		},
		Desc:     "The request could not be understood by the server due to malformed syntax",
		Hint:     "Please check your request parameters and try again",
		Path:     "",
		HttpCode: 400,
		Level:    LevelHighUser,
	}))

	Unauthorized, _ = Registry.Reg(SP(Err{
		Messages: map[string]string{
			En: "Unauthorized",
			Ru: "Не авторизован",
		},
		Desc:     "Authentication is required and has failed or has not been provided",
		Hint:     "Please provide valid authentication credentials",
		Path:     "",
		HttpCode: 401,
		Level:    LevelHighUser,
	}))

	Forbidden, _ = Registry.Reg(SP(Err{
		Messages: map[string]string{
			En: "Forbidden",
			Ru: "Доступ запрещен",
		},
		Desc:     "You don't have permission to access this resource",
		Hint:     "Please contact your administrator if you need access",
		Path:     "",
		HttpCode: 403,
		Level:    LevelHighUser,
	}))

	Timeout, _ = Registry.Reg(SP(Err{
		Messages: map[string]string{
			En: "Request timeout",
			Ru: "Время ожидания истекло",
		},
		Desc:     "The server timed out waiting for the request",
		Hint:     "Please try again. If the problem persists, contact support",
		Path:     "",
		HttpCode: 408,
		Level:    LevelHighUser,
	}))
}

func (r *registry) Reg(e *SPError) (hash.Hash, error) {
	r.Lock()
	defer r.Unlock()

	if e == nil {
		return nil, SP(Err{
			Messages: map[string]string{
				En: "Nil error provided",
				Ru: "передана nil ошибка",
			},
			Desc:     "Provided error is nil. This is not allowed)",
			Hint:     "Please, check your code and provide a valid error",
			Path:     "core/sperr/registry.go:103:1",
			HttpCode: 400,
			Level:    LevelHighDebug,
		})
	}

	h, err := e.Done()
	if err != nil {
		return nil, SP(Err{
			Messages: map[string]string{
				En: "Failed to validate SPError",
				Ru: "Ошибка в процессе валидации SPError",
			},
			Desc:     "Failed to create hash id of your error. It happens when you try to register an error with an empty description. Provided data of error in Meta",
			Hint:     "Please, check your fields and provide a valid description, hint and EN message for your error",
			Path:     "core/sperr/registry.go:103:1",
			HttpCode: 400,
			Level:    LevelHighDebug,
			Cause:    err,
			Meta: map[string]any{
				SPErrorKey: *e,
			},
		})
	}

	r.errs[h] = e
	return h, nil
}

func (r *registry) Get(h hash.Hash) (*SPError, error) {
	r.RLock()
	defer r.RUnlock()

	sp, ok := r.errs[h]
	if !ok {
		return nil, SP(Err{
			Messages: map[string]string{
				En: "No such error",
				Ru: "Такой ошибки не существует",
			},
			Desc:     "There is no error with the provided hash ID",
			Hint:     "Please, check your hash ID or fact of error registration. Registration should be in init() function",
			Path:     "core/sperr/registry.go:144:1",
			HttpCode: 404,
			Level:    LevelHighDebug,
			Meta: map[string]any{
				HashKey: h,
			},
		})
	}

	return sp, nil
}
