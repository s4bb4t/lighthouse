package sperr

import (
	"database/sql"
	"github.com/s4bb4t/lighthouse/core/levels"
	"testing"
)

func TestSPError_DeepIs(t *testing.T) {
	tests := []struct {
		name string
		args error
		err  *SPError
		want bool
	}{
		{
			name: "normal",
			args: sql.ErrNoRows,
			err: Registry.errs[NotFound].Wrap(SP(Err{
				Messages: map[string]string{
					En: "t",
				},
				Desc:  "desc",
				Hint:  "hint",
				Level: levels.LevelHighUser,
				Cause: sql.ErrNoRows,
			})),
			want: true,
		},
		{
			name: "not found",
			args: sql.ErrNoRows,
			err:  Registry.errs[BadRequest].Wrap(Registry.errs[BadRequest]),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.DeepIs(tt.args); got != tt.want {
				t.Errorf("DeepIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestSPError_Is(t *testing.T) {
//	type fields struct {
//		messages          map[string]string
//		desc              string
//		hint              string
//		path              string
//		id                hash.Hash
//		httpCode          int
//		level             levels.ErrorLevel
//		timestamp         time.Time
//		cause             error
//		stack             []string
//		meta              map[string]any
//		remainsUnderlying int
//		underlying        *SPError
//	}
//	type args struct {
//		err error
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			e := &SPError{
//				messages:          tt.fields.messages,
//				desc:              tt.fields.desc,
//				hint:              tt.fields.hint,
//				path:              tt.fields.path,
//				id:                tt.fields.id,
//				httpCode:          tt.fields.httpCode,
//				level:             tt.fields.level,
//				timestamp:         tt.fields.timestamp,
//				cause:             tt.fields.cause,
//				stack:             tt.fields.stack,
//				meta:              tt.fields.meta,
//				remainsUnderlying: tt.fields.remainsUnderlying,
//				underlying:        tt.fields.underlying,
//			}
//			if got := e.Is(tt.args.err); got != tt.want {
//				t.Errorf("Is() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestSPError_IsSP(t *testing.T) {
//	type fields struct {
//		messages          map[string]string
//		desc              string
//		hint              string
//		path              string
//		id                hash.Hash
//		httpCode          int
//		level             levels.ErrorLevel
//		timestamp         time.Time
//		cause             error
//		stack             []string
//		meta              map[string]any
//		remainsUnderlying int
//		underlying        *SPError
//	}
//	type args struct {
//		err *SPError
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			e := &SPError{
//				messages:          tt.fields.messages,
//				desc:              tt.fields.desc,
//				hint:              tt.fields.hint,
//				path:              tt.fields.path,
//				id:                tt.fields.id,
//				httpCode:          tt.fields.httpCode,
//				level:             tt.fields.level,
//				timestamp:         tt.fields.timestamp,
//				cause:             tt.fields.cause,
//				stack:             tt.fields.stack,
//				meta:              tt.fields.meta,
//				remainsUnderlying: tt.fields.remainsUnderlying,
//				underlying:        tt.fields.underlying,
//			}
//			if got := e.IsSP(tt.args.err); got != tt.want {
//				t.Errorf("IsSP() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_cmpHashes(t *testing.T) {
//	type args struct {
//		h1 hash.Hash
//		h2 hash.Hash
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := cmpHashes(tt.args.h1, tt.args.h2); got != tt.want {
//				t.Errorf("cmpHashes() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
