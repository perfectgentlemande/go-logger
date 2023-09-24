package zerolog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/perfectgentlemande/go-logger"
	zlogsentry "github.com/perfectgentlemande/zerolog-sentry"
	"github.com/rs/zerolog"
)

var ErrNoConfig = errors.New("no config")

type zerologWrapper struct {
	log *zerolog.Logger
}

var levelZerolog = map[logger.Level]zerolog.Level{
	logger.PanicLevel: zerolog.PanicLevel,
	logger.FatalLevel: zerolog.FatalLevel,
	logger.ErrorLevel: zerolog.ErrorLevel,
	logger.WarnLevel:  zerolog.WarnLevel,
	logger.InfoLevel:  zerolog.InfoLevel,
	logger.DebugLevel: zerolog.DebugLevel,
}

func extractZerologOutput(value string) io.Writer {
	switch value {
	case logger.OutputStdOut, "":
		return os.Stdout
	case logger.OutputStdErr:
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

func DefaultZerolog() logger.Logger {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).Level(zerolog.DebugLevel)

	return &zerologWrapper{
		log: &log,
	}
}

func prepareSentryWriter(conf *logger.Sentry) (io.Writer, error) {
	if conf == nil {
		return nil, fmt.Errorf("no config for setnry: %w", ErrNoConfig)
	}

	options := []zlogsentry.WriterOption{}
	if conf.StacktraceConfigurationEnable {
		options = append(options, zlogsentry.WithTracing())
	}
	options = append(
		options,
		zlogsentry.WithTags(conf.Tags),
		zlogsentry.WithFlushTimeout(conf.Timeout),
		zlogsentry.WithEnvironment("dev"),
		zlogsentry.WithRelease("1.0.0"),
	)

	zlogsentry.WithLevels()

	w, err := zlogsentry.New(conf.DSN, options...)
	if err != nil {
		return nil, fmt.Errorf("cannot create zlogsentry writer: %w", err)
	}

	return w, nil
}

func NewZerolog(config *logger.Config) logger.Logger {
	output := extractZerologOutput(config.Output)
	log := zerolog.New(output)

	if config.Formatter != logger.FormatterJSON {
		log = zerolog.New(zerolog.ConsoleWriter{Out: output})
	}
	log = log.Level(levelZerolog[config.Level])

	sentryWriter, err := prepareSentryWriter(config.Sentry)
	if err != nil {
		if errors.Is(err, ErrNoConfig) {
			log.Info().Msg("sentry disabled")
		} else {
			log.Error().Err(err).Msg("some error with sentry config")
		}
	} else {
		log = log.Output(zerolog.MultiLevelWriter(output, sentryWriter))
	}

	return &zerologWrapper{
		log: &log,
	}
}

func (zw *zerologWrapper) Panic(msg string) {
	zw.log.Panic().Time(logger.FieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Fatal(msg string) {
	zw.log.Fatal().Time(logger.FieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Error(msg string) {
	zw.log.Error().Time(logger.FieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Warning(msg string) {
	zw.log.Warn().Time(logger.FieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Info(msg string) {
	zw.log.Info().Time(logger.FieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) Debug(msg string) {
	zw.log.Debug().Time(logger.FieldTime, time.Now()).Msg(msg)
}
func (zw *zerologWrapper) WithField(key string, value interface{}) logger.Logger {
	log := zw.log.With().Fields(map[string]interface{}{key: value}).Logger()
	return &zerologWrapper{
		log: &log,
	}
}
func (zw *zerologWrapper) WithFields(fields logger.Fields) logger.Logger {
	log := zw.log.With().Fields(map[string]interface{}(fields)).Logger()
	return &zerologWrapper{
		log: &log,
	}
}
func (zw *zerologWrapper) WithError(err error) logger.Logger {
	return zw.WithField(logger.FieldError, err)
}
