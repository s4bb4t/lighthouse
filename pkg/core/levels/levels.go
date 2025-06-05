package levels

import "strconv"

// Level is a 2^n interpretation of the "weight" of an error.
// It represents how severe the error is and how difficult it is to understand, handle and fix.
type Level uint8

const (
	LevelDebug Level = 255 // the highest debug error level
	LevelError Level = 64
	LevelInfo  Level = 8
	LevelUser  Level = 2
	LevelNoop  Level = 0 // no errors
)

func (e Level) String() string {
	return strconv.Itoa(int(e))
}
