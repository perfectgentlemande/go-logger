package logger

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
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
