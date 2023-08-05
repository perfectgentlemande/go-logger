package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusWrapper struct {
	log *logrus.Entry
}

func newLogrus(config *Config) Logger {
	return &logrusWrapper{}
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
