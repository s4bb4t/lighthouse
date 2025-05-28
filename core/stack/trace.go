package stack

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// TODO: stack trace

func captureTrace() string {
	var trace string
	pcs := make([]uintptr, 10)
	n := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		absPath, _ := filepath.Abs(frame.File)
		trace += fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, absPath, frame.Line)
		if !more {
			break
		}
	}
	return trace
}
