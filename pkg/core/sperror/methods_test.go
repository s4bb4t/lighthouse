package sperror

import (
	"errors"
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"reflect"
	"testing"
)

func TestEnsure(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "regular error",
			args: args{err: New(Sample{
				Messages: map[string]string{En: "Test error"},
				Desc:     "test description",
				Hint:     "test hint",
				Level:    levels.LevelError,
			})},
			want: New(Sample{
				Messages: map[string]string{En: "Test error"},
				Desc:     "test description",
				Hint:     "test hint",
				Level:    levels.LevelError,
			}),
		},
		{
			name: "regular error",
			args: args{err: fmt.Errorf("test cause")},
			want: New(Sample{
				Messages: map[string]string{En: "Unknown error"},
				Desc:     "test cause",
				Hint:     "Check original .Error()",
				Level:    levels.LevelError,
				Cause:    fmt.Errorf("test cause")}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ensure(tt.args.err); !errors.Is(got, tt.want) {
				t.Errorf("Ensure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_AllMeta(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]any
	}{
		{
			name: "empty meta",
			fields: fields{
				meta: nil,
			},
			want: map[string]any{},
		},
		{
			name: "with meta",
			fields: fields{
				meta: map[string]any{"key": "value"},
			},
			want: map[string]any{"key": "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.AllMeta(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Caused(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "no cause",
			fields: fields{
				Core: CoreError{
					Cause: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "with cause",
			fields: fields{
				Core: CoreError{
					Cause: New(Sample{
						Messages: map[string]string{En: "Cause error"},
					}),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if err := e.Caused(); (err != nil) != tt.wantErr {
				t.Errorf("Caused() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestError_Code(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "default code",
			fields: fields{
				User: UserError{
					HttpCode: 500,
				},
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Desc(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "with description",
			fields: fields{
				Core: CoreError{
					Desc: "Test description",
				},
			},
			want: "Test description",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Desc(); got != tt.want {
				t.Errorf("Desc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "error string",
			fields: fields{
				Core: CoreError{
					Desc: "Test error",
					Hint: "Test hint",
				},
			},
			want: "Test error: Test hint",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Hint(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "with hint",
			fields: fields{
				Core: CoreError{
					Hint: "Test hint",
				},
			},
			want: "Test hint",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Hint(); got != tt.want {
				t.Errorf("Hint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Level(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name   string
		fields fields
		want   levels.Level
	}{
		{
			name: "error level",
			fields: fields{
				User: UserError{
					Level: levels.LevelError,
				},
			},
			want: levels.LevelError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Level(); got != tt.want {
				t.Errorf("Level() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Meta(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			name: "nil meta",
			fields: fields{
				meta: nil,
			},
			args: args{key: "test"},
			want: nil,
		},
		{
			name: "existing key",
			fields: fields{
				meta: map[string]any{"test": "value"},
			},
			args: args{key: "test"},
			want: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Meta(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Meta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Msg(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	type args struct {
		lg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "english message",
			fields: fields{
				User: UserError{
					Messages: map[string]string{"en": "Test message"},
				},
			},
			args: args{lg: "en"},
			want: "Test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Msg(tt.args.lg); got != tt.want {
				t.Errorf("Msg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Source(t *testing.T) {
	type fields struct {
		Core              CoreError
		User              UserError
		meta              map[string]any
		remainsUnderlying int
		underlying        *Error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "with source",
			fields: fields{
				Core: CoreError{
					Source: "test_source",
				},
			},
			want: "test_source",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Core:              tt.fields.Core,
				User:              tt.fields.User,
				meta:              tt.fields.meta,
				remainsUnderlying: tt.fields.remainsUnderlying,
				underlying:        tt.fields.underlying,
			}
			if got := e.Source(); got != tt.want {
				t.Errorf("Source() = %v, want %v", got, tt.want)
			}
		})
	}
}
