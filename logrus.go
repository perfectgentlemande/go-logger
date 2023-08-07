package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logrusWrapper struct {
	log *logrus.Entry
}

func extractLogrusOutput(value string) *os.File {
	switch value {
	case OutputStdOut, "":
		return os.Stdout
	case OutputStdErr:
		return os.Stderr
	default:
		var err error
		fl, err := os.OpenFile(value, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			defaultZap().Error("cant create log file, falling to stdout", zap.Field{
				Key:       "error",
				Type:      zapcore.ErrorType,
				Interface: err,
			})
			return os.Stdout
		} else {
			return fl
		}
	}
}

func newLogrus(config *Config) Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	return &logrusWrapper{
		log: logrus.NewEntry(log),
	}
}

func (lw *logrusWrapper) Panic(msg string) {
	lw.log.Panic(msg)
}
func (lw *logrusWrapper) Fatal(msg string) {
	lw.log.Fatal(msg)
}
func (lw *logrusWrapper) Error(msg string) {
	lw.log.Error(msg)
}
func (lw *logrusWrapper) Warning(msg string) {
	lw.log.Warning(msg)
}
func (lw *logrusWrapper) Info(msg string) {
	lw.log.Info(msg)
}
func (lw *logrusWrapper) Debug(msg string) {
	lw.log.Debug(msg)
}
func (lw *logrusWrapper) WithField(key string, value interface{}) Logger {
	return &logrusWrapper{
		log: lw.log.WithField(key, value),
	}
}
func (lw *logrusWrapper) WithFields(fields Fields) Logger {
	return &logrusWrapper{
		log: lw.log.WithFields(logrus.Fields(fields)),
	}
}
func (lw *logrusWrapper) WithError(err error) Logger {
	return &logrusWrapper{
		log: lw.log.WithError(err),
	}
}
