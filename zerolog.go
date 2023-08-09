package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type zerologWrapper struct {
	log    *zerolog.Logger
	fields Fields
}

var levelZerolog = map[Level]zerolog.Level{
	PanicLevel: zerolog.PanicLevel,
	FatalLevel: zerolog.FatalLevel,
	ErrorLevel: zerolog.ErrorLevel,
	WarnLevel:  zerolog.WarnLevel,
	InfoLevel:  zerolog.InfoLevel,
	DebugLevel: zerolog.DebugLevel,
}

func extractZerologOutput(value string) io.Writer {
	switch value {
	case OutputStdOut, "":
		return os.Stdout
	case OutputStdErr:
		return os.Stderr
	default:
		var err error
		fl, err := os.OpenFile(value, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			DefaultZerolog().WithError(err).Error("can't create log file, falling to stdout")
			return os.Stdout
		} else {
			return fl
		}
	}
}

func DefaultZerolog() Logger {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).Level(zerolog.DebugLevel)

	return &zerologWrapper{
		log: &log,
	}
}

func NewZerolog(config *Config) Logger {
	output := extractZerologOutput(config.Output)
	log := zerolog.New(output)

	if config.Formatter != FormatterJSON {
		log = zerolog.New(zerolog.ConsoleWriter{Out: output})
	}
	log = log.Level(levelZerolog[config.Level])

	return &zerologWrapper{
		log: &log,
	}
}

func (zw *zerologWrapper) Panic(msg string) {
	zw.log.Panic().Fields(map[string]interface{}(zw.fields)).Time(fieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Fatal(msg string) {
	zw.log.Fatal().Fields(map[string]interface{}(zw.fields)).Time(fieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Error(msg string) {
	zw.log.Error().Fields(map[string]interface{}(zw.fields)).Time(fieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Warning(msg string) {
	zw.log.Warn().Fields(map[string]interface{}(zw.fields)).Time(fieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Info(msg string) {
	zw.log.Info().Fields(map[string]interface{}(zw.fields)).Time(fieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Debug(msg string) {
	zw.log.Debug().Fields(map[string]interface{}(zw.fields)).Time(fieldTime, time.Now()).Msg(msg)
}

func copyFields(fields Fields) Fields {
	res := make(Fields, len(fields))

	for k := range fields {
		res[k] = fields[k]
	}

	return res
}
func (zw *zerologWrapper) WithField(key string, value interface{}) Logger {
	newFields := copyFields(zw.fields)
	newFields[key] = value

	return &zerologWrapper{
		log:    zw.log,
		fields: newFields,
	}
}
func (zw *zerologWrapper) WithFields(fields Fields) Logger {
	newFields := copyFields(zw.fields)

	for k := range fields {
		newFields[k] = fields[k]
	}

	return &zerologWrapper{
		log:    zw.log,
		fields: newFields,
	}
}
func (zw *zerologWrapper) WithError(err error) Logger {
	return zw.WithField(fieldError, err)
}
