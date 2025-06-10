package sp

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"testing"
)

func ExampleError_Spin() {
	dbErr := New(Sample{
		Messages: map[string]string{
			En: "Database connection error",
		},
		Desc:  "Failed to connect to PostgreSQL",
		Hint:  "Check database URL and credentials",
		Level: levels.LevelDebug,
	})

	appErr := WrapNew(dbErr, Sample{
		Messages: map[string]string{
			En: "Internal application error",
		},
		Desc:  "Could not retrieve user profile",
		Hint:  "Ensure the user ID exists",
		Level: levels.LevelError,
	})

	clientErr := WrapNew(appErr, Sample{
		Messages: map[string]string{
			En: "Something went wrong",
		},
		Desc:  "User profile loading failed",
		Hint:  "Please try again later",
		Level: levels.LevelUser,
	})

	// Show how the same error behaves differently depending on Spin level:

	fmt.Printf("User Level:\n%v\n\n", clientErr.Spin(levels.LevelUser).Desc())
	fmt.Printf("Error Level:\n%v\n\n", clientErr.Spin(levels.LevelError).Desc())
	fmt.Printf("Debug Level:\n%v\n\n", clientErr.Spin(levels.LevelDebug).Desc())

	// Output:
	// User Level:
	// User profile loading failed
	//
	// Error Level:
	// Could not retrieve user profile
	//
	// Debug Level:
	// Failed to connect to PostgreSQL
}
func TestSPError_Spin(t *testing.T) {
	tests := []struct {
		name      string
		err       *Error
		spinLevel levels.Level
		want      string
	}{
		{
			name:      "1",
			err:       Api(),
			spinLevel: levels.LevelError,
			want:      App().Source(),
		},
		{
			name:      "2",
			err:       Api(),
			spinLevel: levels.LevelInfo,
			want:      Api().Source(),
		},
		{
			name:      "3",
			err:       Api(),
			spinLevel: levels.LevelDebug,
			want:      DB().Source(),
		},
		{
			name:      "4",
			err:       Api(),
			spinLevel: levels.LevelUser,
			want:      "/Users/dmitrijbratiskin/GolandProjects/sperr/pkg/core/sp/spin_test.go:103",
		},
		{
			name:      "5",
			err:       Api(),
			spinLevel: 1,
			want:      "/Users/dmitrijbratiskin/GolandProjects/sperr/pkg/core/sp/spin_test.go:103",
		},
		{
			name:      "6",
			err:       NewSpErr(),
			spinLevel: levels.LevelError,
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				return
			}
			got := tt.err.Spin(tt.spinLevel).Source()
			if got != tt.want {
				t.Errorf("Error.Spin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Api() *Error {
	err := App()
	return WrapNew(err, Sample{
		Messages: map[string]string{
			En: "Internal",
		},
		Desc:  "Internal Error",
		Hint:  "Try Again later",
		Level: levels.LevelInfo,
	})
}

func App() *Error {
	err := DB()
	return WrapNew(err, Sample{
		Messages: map[string]string{
			En: "App err",
		},
		Desc:  "Database error",
		Hint:  "Check repo layer",
		Level: levels.LevelError,
	})
}

func DB() *Error {
	return New(Sample{
		Messages: map[string]string{
			En: "Db connection failed",
		},
		Desc:  "Failed to connect to storage",
		Hint:  "check connection string, credentials, etc.",
		Level: levels.LevelDebug,
	})
}
