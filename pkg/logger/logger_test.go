package logger

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	sp2 "github.com/s4bb4t/lighthouse/pkg/core/sp"
	"os"
	"testing"
)

var l = New(Local, "en", os.Stdout)
var d = New(Dev, "en", os.Stdout)

func TestLogger_Dev(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args *sp2.Error
	}{
		{
			name: "normal",
			args: sp2.New(sp2.Err{
				Messages: map[string]string{
					"en": "Failed to connect to storage",
				},
				Desc:     "Failed to connect to storage due to timeout",
				Hint:     "Check connection string, credentials, etc.",
				HttpCode: 500,
				Level:    levels.LevelHighUser,
			}).MustDone(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d.Error(tt.args, levels.LevelDeepDebug)
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				msg: "test",
				args: []any{
					"key",
					"val",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Debug(tt.args.msg, tt.args.args...)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type args struct {
		e   error
		lvl levels.Level
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				e: sp2.New(sp2.Err{
					Messages: map[string]string{
						"en": "test",
					},
					Desc: "123",
					Hint: "123",
				}).MustDone(),
				lvl: levels.LevelError,
			},
		},
		{
			name: "basic",
			args: args{
				e:   os.ErrClosed,
				lvl: levels.LevelError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Error(tt.args.e, tt.args.lvl)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				msg: "hi",
			},
		},
		{
			name: "normal with args",
			args: args{
				msg: "hi with args",
				args: []any{
					"key",
					"val",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Info(tt.args.msg, tt.args.args...)
		})
	}
}

//func TestLogger_With(t *testing.T) {
//	type args struct {
//		args []any
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "normal",
//			args: args{args: []any{
//				"key", "val",
//			}},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l.With(tt.args.args...)
//			l.Info("hi")
//		})
//	}
//}
