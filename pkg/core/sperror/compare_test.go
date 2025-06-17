package sperror

import (
	"database/sql"
	"errors"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"testing"
)

func TestError_Is(t *testing.T) {
	var ErrTemporaryUnavailable = New(Sample{
		Messages: map[string]string{
			En: "Temporary unavailable - initializing",
			Ru: "Временно недоступно - инициализация",
		},
		Desc:     "We are in process of initialization",
		Hint:     "Please try again later",
		HttpCode: 307,
		Level:    levels.LevelUser,
	})

	if !ErrTemporaryUnavailable.Is(ErrTemporaryUnavailable) {
		t.Fail()
	}
	if !errors.Is(ErrTemporaryUnavailable, ErrTemporaryUnavailable) {
		t.Fail()
	}
	if !errors.Is(ErrTemporaryUnavailable, Any(ErrTemporaryUnavailable, "hi", "there")) {
		t.Fail()
	}

	err := Any(ErrTemporaryUnavailable, "test", "test")
	for i := 0; i < 100; i++ {
		err = Any(err, "test", "test")
	}
	if !errors.Is(err, ErrTemporaryUnavailable) {
		t.Fail()
	}

	err2 := Any(sql.ErrNoRows, "test", "test")
	err1 := Any(sql.ErrNoRows, "test2", "test2")
	if !errors.Is(err2, err1) {
		t.Fail()
	}
}
