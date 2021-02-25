package logger

var logger logList

func Add(log Log) {
	logger = append(logger, log)
}

func Fatal(v interface{}) {
	logger.log(LevelFatal, v)
}

func Error(v interface{}) {
	logger.log(LevelError, v)
}

func Warn(v interface{}) {
	logger.log(LevelWarn, v)
}

func Info(v interface{}) {
	logger.log(LevelInfo, v)
}

func Debug(v interface{}) {
	logger.log(LevelDebug, v)
}
