package sp

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"hash"
	"reflect"
	"testing"
)

func Test_registry_Get(t *testing.T) {
	tests := []struct {
		name    string
		h       hash.Hash
		want    *Error
		wantErr bool
	}{
		{
			name: "Normal - Internal",
			h:    Internal,
			want: New(Err{
				Messages: map[string]string{
					En: "Internal server error",
					Ru: "Ошибка сервера",
				},
				Desc:     "Internal server error. We are sorry for the inconvenience.",
				Hint:     "Please try again later - we are working on it.",
				HttpCode: 500,
				Level:    levels.LevelUser,
			}),
			wantErr: false,
		},
		{
			name:    "Error",
			h:       nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Registry.Get(tt.h)
			if !got.IsSP(tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			t.Log("\n", got.ReadDesc(), "\n", tt.want.ReadDesc())
		})
	}
}

func Test_registry_Reg(t *testing.T) {
	tests := []struct {
		name    string
		err     *Error
		want    hash.Hash
		wantErr bool
	}{
		{
			name:    "Exists",
			err:     Registry.errs[Internal],
			want:    Internal,
			wantErr: false,
		},
		{
			name: "Normal",
			err: New(Err{
				Messages: map[string]string{
					En: "test",
					Ru: "test",
				},
				Desc:     "test.",
				Hint:     "test.",
				HttpCode: 200,
				Level:    levels.LevelUser,
			}),
			wantErr: false,
		},
		{
			name:    "Error",
			err:     nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Registry.Reg(tt.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Reg() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
