package export

import (
	"github.com/s4bb4t/lighthouse/core/sp"
	"testing"
)

func TestExport(t *testing.T) {
	tests := []struct {
		name    string
		e       *sp.Error
		exp     func(e ...*sp.Error) ([]byte, error)
		wantErr bool
	}{
		{
			name:    "csv",
			e:       sp.Registry.Get(sp.Internal),
			exp:     CSV,
			wantErr: false,
		},
		{
			name:    "xml",
			e:       sp.Registry.Get(sp.Internal),
			exp:     XML,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.exp(tt.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(string(got))
		})
	}
}
