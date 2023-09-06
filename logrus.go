package logger

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/makasim/sentryhook"
	"github.com/sirupsen/logrus"
)

type logrusWrapper struct {
	log *logrus.Entry
}

var levelLogrus = map[Level]logrus.Level{
	PanicLevel: logrus.PanicLevel,
	FatalLevel: logrus.FatalLevel,
	ErrorLevel: logrus.ErrorLevel,
	WarnLevel:  logrus.WarnLevel,
	InfoLevel:  logrus.InfoLevel,
	DebugLevel: logrus.DebugLevel,
}

// logrusLevelMap
var logrusLevelMap = map[Level]logrus.Level{
	emptyLevel: logrus.TraceLevel,
	PanicLevel: logrus.PanicLevel,
	FatalLevel: logrus.FatalLevel,
	ErrorLevel: logrus.ErrorLevel,
	WarnLevel:  logrus.WarnLevel,
	InfoLevel:  logrus.InfoLevel,
	DebugLevel: logrus.DebugLevel,
}

func DefaultLogrus() Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logrus.DebugLevel)

	return &logrusWrapper{
		log: logrus.NewEntry(log),
	}
}

func extractLogrusOutput(value string) *os.File {
	switch value {
	case OutputStdOut, "":
		return os.Stdout
	case OutputStdErr:
		return os.Stderr
	default:
		var err error
		fl, err := os.OpenFile(value, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			DefaultLogrus().WithError(err).Error("cant create log file, falling to stdout")
			return os.Stdout
		} else {
			return fl
		}
	}
}

func sentryLogrusHook(config *Sentry) (*sentryhook.Hook, error) {
	tr := sentry.NewHTTPTransport()
	if config.Timeout == 0 {
		tr.Timeout = 2 * time.Second
	} else {
		tr.Timeout = config.Timeout
	}

	opts := sentry.ClientOptions{
		Dsn:              config.DSN,
		Transport:        tr,
		AttachStacktrace: config.StacktraceConfigurationEnable,
	}

	if config.SSLSkipVerify {
		opts.HTTPTransport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	if len(config.Tags) != 0 {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTags(config.Tags)
		})
	}

	if err := sentry.Init(opts); err != nil {
		return nil, err
	}

	hook := sentryhook.New(
		logLevels(config),
	)

	return &hook, nil
}

func addSentryLogrusHook(logger *logrus.Logger, config *Config) {
	if config.Sentry != nil {
		sh, err := sentryLogrusHook(config.Sentry)
		if err == nil {
			logger.AddHook(sh)
		} else {
			logger.WithError(err).Debug("can't add hook sentry")
		}
	}
}

func NewLogrus(config *Config) Logger {
	log := logrus.New()
	log.SetOutput(extractLogrusOutput(config.Output))

	if config.Formatter == FormatterJSON {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{})
	}
	log.SetLevel(levelLogrus[config.Level])
	addSentryLogrusHook(log, config)

	return &logrusWrapper{
		log: logrus.NewEntry(log),
	}
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
