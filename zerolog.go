package logger

import (
	"github.com/rs/zerolog"
)

type zerologWrapper struct {
	log *zerolog.Logger
}

func newZerolog(config *Config) Logger {
	return &zerologWrapper{}
}

func (zw *zerologWrapper) Debug(msg string) {
	zw.log.Debug().Msg(msg)
}
func (zw *zerologWrapper) Info(msg string) {
	zw.log.Info().Msg(msg)
}
func (zw *zerologWrapper) Warning(msg string) {
	zw.log.Warn().Msg(msg)
}
func (zw *zerologWrapper) Error(msg string) {
	zw.log.Error().Msg(msg)
}
func (zw *zerologWrapper) Fatal(msg string) {
	zw.log.Fatal().Msg(msg)
}
