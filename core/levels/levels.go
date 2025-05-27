package levels

type ErrorLevel uint8

const (
	LevelDeepDebug   ErrorLevel = 255 // the highest debug error level
	LevelMediumDebug ErrorLevel = 128
	LevelHighDebug   ErrorLevel = 64
	LevelError       ErrorLevel = 32
	LevelWarn        ErrorLevel = 16
	LevelInfo        ErrorLevel = 8
	LevelLowUser     ErrorLevel = 4
	LevelMediumUser  ErrorLevel = 2
	LevelHighUser    ErrorLevel = 1
	LevelNoop        ErrorLevel = 0 // no errors
)
