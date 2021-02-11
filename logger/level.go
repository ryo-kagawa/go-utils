package logger

type Level uint

const (
	LevelUnknown Level = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

func ParseLevel(v string) Level {
	switch v {
	case "FATAL":
		return LevelFatal
	case "ERROR":
		return LevelError
	case "WARN":
		return LevelWarn
	case "INFO":
		return LevelInfo
	case "DEBUG":
		return LevelDebug
	}
	return LevelUnknown
}
