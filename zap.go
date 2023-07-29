package logger

import (
	"go.uber.org/zap"
)

type zapWrapper struct {
	log *zap.Logger
}

func newZap(config *Config) Logger {
	return &zapWrapper{}
}

func (zw *zapWrapper) Debug(args ...interface{}) {

}
func (zw *zapWrapper) Info(args ...interface{}) {

}
func (zw *zapWrapper) Warning(args ...interface{}) {}
func (zw *zapWrapper) Error(args ...interface{}) {

}
func (zw *zapWrapper) Fatal(args ...interface{}) {

}
