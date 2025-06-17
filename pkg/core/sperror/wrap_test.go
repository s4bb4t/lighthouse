package sperror

import (
	"fmt"
	"testing"
)

func ExampleError_Unwrap() {
	inner := New(Sample{
		Desc:  "Inner error description",
		Hint:  "Try again later",
		Cause: fmt.Errorf("example error"),
	})

	outer := New(Sample{}).Wrap(inner)

	err := outer.Unwrap()
	fmt.Printf("%v\nType: %T\n\n", err, err)
	e := Ensure(err).Unwrap()
	fmt.Printf("%v", e)

	// Output:
	//
	// Inner error description: Try again later
	// Type: *sperror.Error
	//
	// example error
}

func ExampleError_Wrap() {
	inner := New(Sample{
		Desc: "Inner error",
		Hint: "hint",
	})

	outer := New(Sample{
		Desc: "Application error description",
		Hint: "hint",
	}).Wrap(inner)

	fmt.Printf("%v\n", outer.Unwrap())
	// Output:
	// Inner error: hint
}

func ExampleWrap() {
	src := New(Sample{})

	dst := New(Sample{
		Desc: "Application error description",
		Hint: "hint",
	})

	wrapped := Wrap(src, dst)
	fmt.Printf("%v\n", wrapped)
	// Output:
	// Application error description: hint
}

func ExampleWrapNew() {
	src := New(Sample{
		Desc: "Application error description",
		Hint: "hint",
	})

	wrapped := WrapNew(src, Sample{})

	fmt.Printf("%v\n", wrapped.Unwrap())
	// Output: Application error description: hint
}

func TestError_Unwrap(t *testing.T) {
	tests := []struct {
		name string
		err  *Error
		want error
	}{
		{
			name: "unwrap with underlying error",
			err: New(Sample{}).Wrap(New(Sample{
				Desc:  "Inner error",
				Hint:  "hint",
				Cause: fmt.Errorf("cause"),
			})),
			want: New(Sample{
				Desc:  "Inner error",
				Hint:  "hint",
				Cause: fmt.Errorf("cause"),
			}),
		},
		{
			name: "unwrap with cause only",
			err: New(Sample{
				Cause: fmt.Errorf("cause"),
			}),
			want: fmt.Errorf("cause"),
		},
		{
			name: "unwrap with desc and hint only",
			err: New(Sample{
				Desc: "Inner error",
				Hint: "hint",
			}),
			want: fmt.Errorf("Inner error: hint"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Unwrap(); got.Error() != tt.want.Error() {
				t.Errorf("Error.Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Wrap(t *testing.T) {
	tests := []struct {
		name string
		err  *Error
		src  *Error
		want string
	}{
		{
			name: "wrap error",
			err:  New(Sample{Desc: "outer"}),
			src:  New(Sample{Desc: "inner"}),
			want: "outer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Wrap(tt.src); got.Desc() != tt.want {
				t.Errorf("Error.Wrap() = %v, want %v", got.Desc(), tt.want)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name string
		src  *Error
		dst  *Error
		want string
	}{
		{
			name: "wrap error",
			src:  New(Sample{Desc: "source"}),
			dst:  New(Sample{Desc: "destination"}),
			want: "destination",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.src, tt.dst); got.Desc() != tt.want {
				t.Errorf("Wrap() = %v, want %v", got.Desc(), tt.want)
			}
		})
	}
}

func TestWrapNew(t *testing.T) {
	tests := []struct {
		name string
		src  *Error
		dst  Sample
		want string
	}{
		{
			name: "wrap new error",
			src:  New(Sample{Desc: "source"}),
			dst:  Sample{Desc: "destination"},
			want: "destination",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WrapNew(tt.src, tt.dst); got.Desc() != tt.want {
				t.Errorf("WrapNew() = %v, want %v", got.Desc(), tt.want)
			}
		})
	}
}
