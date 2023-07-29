package logger

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type LoggerName string

const (
	Zap     LoggerName = "zap"
	Logrus  LoggerName = "logrus"
	Zerolog LoggerName = "zerolog"
)

func CreateLogger(config *Config) Logger {
	return newZap(config)
}
