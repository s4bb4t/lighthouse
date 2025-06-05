package usecase

import "github.com/s4bb4t/lighthouse/pkg/core/levels"

type (
	Notify interface {
		Info(msg string) error
		Error(err error, group ...string) error
		//Warn(string) error
		//Debug(string) error
	}

	Logger interface {
		Error(e error, lvl levels.Level)
		Debug(msg string, args ...any)
		Info(msg string, args ...any)
	}

	Storage interface {
		Put(group string, id int64) error
		Read(group string) (ids []int64, err error)
	}
)
