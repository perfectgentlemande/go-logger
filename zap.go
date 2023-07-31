package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapWrapper struct {
	log *zap.Logger
}

var levelZap = map[Level]zapcore.Level{
	PanicLevel: zapcore.PanicLevel,
	FatalLevel: zapcore.FatalLevel,
	ErrorLevel: zapcore.ErrorLevel,
	WarnLevel:  zapcore.WarnLevel,
	InfoLevel:  zapcore.InfoLevel,
	DebugLevel: zapcore.DebugLevel,
}

func newZap(config *Config) Logger {
	log := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), os.Stdout, levelZap[config.Level]))

	return &zapWrapper{
		log: log,
	}
}

func (zw *zapWrapper) Debug(msg string) {
	zw.log.Debug(msg)
}
func (zw *zapWrapper) Info(msg string) {
	zw.log.Info(msg)
}
func (zw *zapWrapper) Warning(msg string) {
	zw.log.Warn(msg)
}
func (zw *zapWrapper) Error(msg string) {
	zw.log.Error(msg)
}
func (zw *zapWrapper) Fatal(msg string) {
	zw.log.Fatal(msg)
}
