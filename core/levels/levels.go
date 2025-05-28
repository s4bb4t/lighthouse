package levels

// Level is a 2^n interpretation of the "weight" of an error.
// It represents how severe the error is and how difficult it is to understand, handle and fix.
type Level uint8

const (
	LevelDeepDebug   Level = 255 // the highest debug error level
	LevelMediumDebug Level = 128
	LevelHighDebug   Level = 64
	LevelError       Level = 32
	LevelWarn        Level = 16
	LevelInfo        Level = 8
	LevelLowUser     Level = 4
	LevelMediumUser  Level = 2
	LevelHighUser    Level = 1
	LevelNoop        Level = 0 // no errors
)
