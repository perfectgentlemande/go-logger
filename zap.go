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

func extractZapOutput(value string) *os.File {
	switch value {
	case OutputStdOut, "":
		return os.Stdout
	case OutputStdErr:
		return os.Stderr
	default:
		var err error
		fl, err := os.OpenFile(value, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			DefaultZap().WithError(err).Error("cant create log file, falling to stdout")
			return os.Stdout
		} else {
			return fl
		}
	}
}

func DefaultZap() Logger {
	return &zapWrapper{
		log: zap.New(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()), os.Stdout, zap.DebugLevel)),
	}
}

func NewZap(config *Config) Logger {
	var encoder zapcore.Encoder

	if config.Formatter == FormatterJSON {
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	}

	log := zap.New(
		zapcore.NewCore(encoder, extractZapOutput(config.Output), levelZap[config.Level]))

	return &zapWrapper{
		log: log,
	}
}

func (zw *zapWrapper) Panic(msg string) {
	zw.log.Panic(msg)
}
func (zw *zapWrapper) Fatal(msg string) {
	zw.log.Fatal(msg)
}
func (zw *zapWrapper) Error(msg string) {
	zw.log.Error(msg)
}
func (zw *zapWrapper) Warning(msg string) {
	zw.log.Warn(msg)
}
func (zw *zapWrapper) Info(msg string) {
	zw.log.Info(msg)
}
func (zw *zapWrapper) Debug(msg string) {
	zw.log.Debug(msg)
}
func (zw *zapWrapper) WithField(key string, value interface{}) Logger {
	field := zapcore.Field{
		Key:       key,
		Type:      zapcore.ReflectType,
		Interface: value,
	}

	return &zapWrapper{
		log: zw.log.With(field),
	}
}
func (zw *zapWrapper) WithFields(fields Fields) Logger {
	res := make([]zapcore.Field, 0, len(fields))

	for key := range fields {
		res = append(res, zapcore.Field{
			Key:       key,
			Type:      zapcore.ReflectType,
			Interface: fields[key],
		})
	}

	return &zapWrapper{
		log: zw.log.With(res...),
	}
}
func (zw *zapWrapper) WithError(err error) Logger {
	field := zapcore.Field{
		Key:       fieldError,
		Type:      zapcore.ErrorType,
		Interface: err,
	}

	return &zapWrapper{
		log: zw.log.With(field),
	}
}
