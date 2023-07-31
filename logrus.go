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

func (lw *logrusWrapper) Debug(msg string) {
	lw.log.Debug(msg)
}
func (lw *logrusWrapper) Info(msg string) {
	lw.log.Info(msg)
}
func (lw *logrusWrapper) Warning(msg string) {
	lw.log.Warning(msg)
}
func (lw *logrusWrapper) Error(msg string) {
	lw.log.Error(msg)
}
func (lw *logrusWrapper) Fatal(msg string) {
	lw.log.Fatal(msg)
}
