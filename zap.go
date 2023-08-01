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

func extractOutput(value string) *os.File {
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

func defaultZap() *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), os.Stdout, zap.ErrorLevel))
}

func newZap(config *Config) Logger {
	log := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), extractOutput(config.Output), levelZap[config.Level]))

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
