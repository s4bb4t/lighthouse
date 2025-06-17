package sperror

import (
	"database/sql"
	"errors"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"testing"
)

//
//func TestSPError_DeepIs(t *testing.T) {
//	tests := []struct {
//		name string
//		args error
//		err  *Error
//		want bool
//	}{
//		{
//			name: "normal",
//			args: sql.ErrNoRows,
//			err: Registry.errs[NotFound].Wrap(New(Sample{
//				Messages: map[string]string{
//					En: "t",
//				},
//				Desc:  "desc",
//				Hint:  "hint",
//				Level: levels.LevelUser,
//				Cause: sql.ErrNoRows,
//			})),
//			want: true,
//		},
//		{
//			name: "not found",
//			args: sql.ErrNoRows,
//			err:  Registry.errs[BadRequest].Wrap(Registry.errs[BadRequest]),
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.err.DeepIs(tt.args); got != tt.want {
//				t.Errorf("DeepIs() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestSPError_Is(t *testing.T) {
//	tests := []struct {
//		name string
//		err  error
//		args *Error
//		want bool
//	}{
//		{
//			name: "normal",
//			args: New(Sample{
//				Messages: map[string]string{
//					En: "en",
//				},
//				Desc:  "1",
//				Hint:  "1",
//				Level: levels.LevelError,
//				Cause: sql.ErrNoRows,
//			}),
//			err:  sql.ErrNoRows,
//			want: true,
//		},
//		{
//			name: "notfound",
//			args: New(Sample{
//				Messages: map[string]string{
//					En: "en",
//				},
//				Desc:  "1",
//				Hint:  "1",
//				Level: levels.LevelError,
//				Cause: sql.ErrTxDone,
//			}),
//			err:  sql.ErrNoRows,
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.args.Is(tt.err); got != tt.want {
//				t.Errorf("Is() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestSPError_IsSP(t *testing.T) {
//	tests := []struct {
//		name string
//		sp1  *Error
//		sp2  *Error
//		want bool
//	}{
//		{
//			name: "true",
//			sp1: New(Sample{
//				Messages: map[string]string{
//					En: "en",
//				},
//				Desc:  "1",
//				Hint:  "1",
//				Level: levels.LevelError,
//				Cause: sql.ErrTxDone,
//			}),
//			sp2: New(Sample{
//				Messages: map[string]string{
//					En: "en",
//				},
//				Desc:  "1",
//				Hint:  "1",
//				Level: levels.LevelError,
//				Cause: sql.ErrTxDone,
//			}),
//			want: true,
//		},
//		{
//			name: "true",
//			sp1: New(Sample{
//				Messages: map[string]string{
//					En: "en",
//				},
//				Desc:  "1",
//				Hint:  "1",
//				Level: levels.LevelError,
//				Cause: sql.ErrTxDone,
//			}),
//			sp2: New(Sample{
//				Messages: map[string]string{
//					En: "en",
//				},
//				Desc:  "2",
//				Hint:  "2",
//				Level: levels.LevelError,
//				Cause: sql.ErrTxDone,
//			}),
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.sp1.IsSP(tt.sp2); got != tt.want {
//				t.Errorf("IsSP() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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

func TestError_Is(t *testing.T) {
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
