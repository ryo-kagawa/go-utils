package logger

type Log interface {
	IsLog(level Level) bool
	Log(v string) error
}
