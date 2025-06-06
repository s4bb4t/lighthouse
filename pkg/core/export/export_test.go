package export

// TODO: refactor tests
//
//func TestExport(t *testing.T) {
//	tests := []struct {
//		name    string
//		e       *sp2.Error
//		exp     func(e ...*sp2.Error) ([]byte, error)
//		wantErr bool
//	}{
//		{
//			name:    "csv",
//			e:       sp2.Registry.Get(sp2.Internal),
//			exp:     CSV,
//			wantErr: false,
//		},
//		{
//			name:    "xml",
//			e:       sp2.Registry.Get(sp2.Internal),
//			exp:     XML,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.exp(tt.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CSV() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			t.Log(string(got))
//		})
//	}
//}
