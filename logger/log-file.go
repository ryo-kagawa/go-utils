package logger

import "os"

type LogFile struct {
	Level Level
	File  *os.File
}

var _ Log = &LogFile{}

func NewLogFile(level Level, file *os.File) *LogFile {
	return &LogFile{
		Level: level,
		File:  file,
	}
}

func (l *LogFile) IsLog(level Level) bool {
	return uint(l.Level) >= uint(level)
}

func (l *LogFile) Log(v string) error {
	_, err := l.File.Write([]byte(v))
	return err
}

func (l *LogFile) Close() error {
	return l.File.Close()
}
