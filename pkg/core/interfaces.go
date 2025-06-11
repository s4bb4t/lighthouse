package core

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
)

type (

	// Notify defines methods for sending informational messages and handling errors with optional grouping.
	// Info sends an informational message and returns an error if it fails.
	// Error logs an error with optional group categorization and returns an error if it fails.
	Notify interface {
		Info(msg string) error
		Error(err error, group ...string) error
	}

	// Logger defines methods for logging errors and informational messages.
	Logger interface {
		ErrorWithLevel(e error, lvl levels.Level)
		Error(e error)
		Warn(msg string, e error, args ...any)
		Debug(msg string, args ...any)
		Info(msg string, args ...any)
	}

	// Storage defines methods for storing and retrieving users groups.
	Storage interface {
		Put(group string, id int64) error
		Read(group string) (ids []int64, err error)
	}

	// Registry defines methods for storing and retrieving pre-defined errors.
	Registry interface {
		Get(id int) error
		Reg(err error)
	}
)
