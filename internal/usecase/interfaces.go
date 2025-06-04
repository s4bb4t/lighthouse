package usecase

import "github.com/s4bb4t/lighthouse/pkg/core/levels"

type (
	Notify interface {
		Info(string) error
		Warn(string) error
		Error(error) error
		Debug(string) error
	}

	Logger interface {
		Error(e error, lvl levels.Level)
		Debug(msg string, args ...any)
		Info(msg string, args ...any)
	}

	Storage interface {
		Put(string, int64) error
		Read(string) ([]int64, error)
	}
)
