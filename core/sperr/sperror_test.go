package sperr

import (
	"testing"
)

func TestSPError_Done(t *testing.T) {
	tests := []struct {
		name      string
		spError   *SPError
		expectErr bool
	}{
		{
			name: "valid hash write",
			spError: &SPError{
				desc: "description",
				hint: "hint",
				messages: map[string]string{
					En: "error message",
				},
			},
			expectErr: false,
		},
		{
			name: "valid hash write",
			spError: &SPError{
				desc: "description1",
				hint: "hint",
				messages: map[string]string{
					En: "error message",
				},
			},
			expectErr: false,
		},
		{
			name:      "empty",
			spError:   &SPError{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.spError.Done()
			if (err != nil) != tt.expectErr {
				t.Errorf("Done() error = %v, expectErr %v", err, tt.expectErr)
			} else if err == nil {
				t.Log(string(tt.spError.id.Sum(nil)))
			}
		})
	}
}

func TestSPError_IsSP(t *testing.T) {
	tests := []struct {
		name     string
		spError1 *SPError
		spError2 *SPError
		wantBool bool
	}{
		{
			name: "same hash write",
			spError1: &SPError{
				messages: map[string]string{
					En: "error message",
				},
				desc: "description",
				hint: "hint",
			},
			spError2: &SPError{
				messages: map[string]string{
					En: "error message",
				},
				desc: "description",
				hint: "hint",
			},
			wantBool: true,
		},
		{
			name: "valid hash write",
			spError1: &SPError{
				messages: map[string]string{
					En: "error message",
				},
				desc: "description",
				hint: "hint",
			},
			spError2: &SPError{
				messages: map[string]string{
					En: "error message1",
				},
				desc: "description1",
				hint: "hint",
			},
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err1 := tt.spError1.Done()
			_, err2 := tt.spError2.Done()
			if err1 != nil || err2 != nil {
				t.Fatalf("errors: 1: %v, 2: %v", err1, err2)
			}
			t.Log(tt.spError1.IsSP(tt.spError2))
			if tt.spError1.IsSP(tt.spError2) != tt.wantBool {
				t.Log(string(tt.spError1.id.Sum(nil)), string(tt.spError2.id.Sum(nil)))
				t.Fail()
			}
		})
	}
}
