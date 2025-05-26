package sperr

type ErrorLevel uint8

const (
	SPErrorKey = "sperror"
	HashKey    = "hash"

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

	En = "en"
	Ru = "ru"
	De = "de"
	Fr = "fr"
	Es = "es"
	Pt = "pt"
	It = "it"
	Nl = "nl"
	Pl = "pl"
	Uk = "uk"
	Cz = "cz"
	Tr = "tr"
	Ar = "ar"
	Ja = "ja"
	Ko = "ko"
	Zh = "zh"
	No = "no"
	Sv = "sv"
	Fi = "fi"
	Da = "da"
	Is = "is"
	Cs = "cs"
	El = "el"
	Hu = "hu"
	Ro = "ro"
	Bg = "bg"
	Lt = "lt"
	Sk = "sk"
	Sl = "sl"
	Hr = "hr"
	Th = "th"
	Lv = "lv"
	Et = "et"
	He = "he"
)
