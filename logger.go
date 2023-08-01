package logger

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger
}

// Fields is a log fields type
type Fields map[string]interface{}

type LoggerName string

const (
	fieldError = "error"

	FormatterJSON = "json"
	FormatterText = "text"

	OutputStdOut = "stdout"
	OutputStdErr = "stderr"

	Zap     LoggerName = "zap"
	Logrus  LoggerName = "logrus"
	Zerolog LoggerName = "zerolog"
)

func CreateLogger(config *Config) Logger {
	return newZap(config)
}
