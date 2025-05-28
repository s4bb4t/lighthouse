package sp

import (
	"github.com/s4bb4t/lighthouse/core/levels"
	"testing"

	"github.com/mailru/easyjson/jwriter"
)

func TestMarshalEasyJSON(t *testing.T) {
	tests := []struct {
		name    string
		spErr   *SPError
		want    string
		wantErr bool
	}{
		{
			name: "nil_fields",
			spErr: &SPError{
				messages: nil,
				meta:     nil,
			},
			want:    `{"messages":null,"description":"","hint":"","http_code":0,"level":0,"meta":null}`,
			wantErr: false,
		},
		{
			name: "populated_fields",
			spErr: &SPError{
				messages: map[string]string{"en": "error message", "fr": "message d'erreur"},
				desc:     "Test description",
				hint:     "Test hint",
				httpCode: 404,
				level:    levels.LevelInfo,
				meta:     map[string]any{"key1": "value1", "key2": 123},
			},
			want:    `{"messages":{"en":"error message","fr":"message d'erreur"},"description":"Test description","hint":"Test hint","http_code":404,"level":8,"meta":{"key1":"value1","key2":123}}`,
			wantErr: false,
		},
		{
			name: "empty_messages",
			spErr: &SPError{
				messages: map[string]string{},
				desc:     "Empty messages",
				meta:     map[string]any{"key": "value"},
			},
			want:    `{"messages":{},"description":"Empty messages","hint":"","http_code":0,"level":0,"meta":{"key":"value"}}`,
			wantErr: false,
		},
		{
			name: "meta_is_empty_map",
			spErr: &SPError{
				messages: map[string]string{"en": "error"},
				meta:     map[string]any{},
			},
			want:    `{"messages":{"en":"error"},"description":"","hint":"","http_code":0,"level":0,"meta":{}}`,
			wantErr: false,
		},
		{
			name: "special_characters_in_messages",
			spErr: &SPError{
				messages: map[string]string{"en": "Error \"message\" with special JSON: {} [] chars"},
				desc:     "Special characters description",
				hint:     "Special characters hint",
			},
			want:    `{"messages":{"en":"Error \"message\" with special JSON: {} [] chars"},"description":"Special characters description","hint":"Special characters hint","http_code":0,"level":0,"meta":null}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w jwriter.Writer
			tt.spErr.MarshalEasyJSON(&w)

			if (w.Error != nil) != tt.wantErr {
				t.Errorf("MarshalEasyJSON() error = %v, wantErr %v", w.Error, tt.wantErr)
				return
			}

			got := string(w.Buffer.BuildBytes())
			if got != tt.want {
				t.Errorf("MarshalEasyJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
