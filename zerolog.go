package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type zerologWrapper struct {
	log *zerolog.Logger
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
			defaultZerolog().Error().Err(err).Msg("can't create log file, falling to stdout")
			return os.Stdout
		} else {
			return fl
		}
	}
}

func defaultZerolog() *zerolog.Logger {
	log := zerolog.New(os.Stdout)

	return &log
}

func newZerolog(config *Config) Logger {
	log := zerolog.New(extractZerologOutput(config.Output))

	return &zerologWrapper{
		log: &log,
	}
}

func (zw *zerologWrapper) Panic(msg string) {
	zw.log.Panic().Msg(msg)
}
func (zw *zerologWrapper) Fatal(msg string) {
	zw.log.Fatal().Msg(msg)
}
func (zw *zerologWrapper) Error(msg string) {
	zw.log.Error().Msg(msg)
}
func (zw *zerologWrapper) Warning(msg string) {
	zw.log.Warn().Msg(msg)
}
func (zw *zerologWrapper) Info(msg string) {
	zw.log.Info().Msg(msg)
}
func (zw *zerologWrapper) Debug(msg string) {
	zw.log.Debug().Msg(msg)
}
func (zw *zerologWrapper) WithField(key string, value interface{}) Logger {

}
func (zw *zerologWrapper) WithFields(fields Fields) Logger {

}
func (zw *zerologWrapper) WithError(err error) Logger {

}
