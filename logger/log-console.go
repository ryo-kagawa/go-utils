package logger

import "io"

type LogConsole struct {
	Level  Level
	Writer io.Writer
}

var _ Log = &LogConsole{}

func NewLogConsole(level Level, writer io.Writer) *LogConsole {
	return &LogConsole{
		Level:  level,
		Writer: writer,
	}
}

func (l *LogConsole) IsLog(level Level) bool {
	return uint(l.Level) >= uint(level)
}

func (l *LogConsole) Log(v string) error {
	_, err := l.Writer.Write([]byte(v))
	return err
}
