package logger

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/getsentry/sentry-go"
	"github.com/makasim/sentryhook"
)

type Sentry struct {
	DSN                           string            `yaml:"dsn" json:"dsn"`
	Level                         Level             `yaml:"level" json:"level"`
	Tags                          map[string]string `yaml:"tags" json:"tags"`
	Timeout                       time.Duration     `yaml:"timeout" json:"timeout"`
	StacktraceConfigurationEnable bool              `yaml:"stacktrace_enable" json:"stacktrace_enable"`
	SSLSkipVerify                 bool              `yaml:"ssl_skip_verify" json:"ssl_skip_verify"`
}

func logLevels(config *Sentry) []logrus.Level {
	levels := make([]logrus.Level, 0, config.Level+1)
	for i := PanicLevel; i <= config.Level; i++ {
		levels = append(levels, logrusLevelMap[i])
	}
	return levels
}

func sentryHook(config *Sentry) (*sentryhook.Hook, error) {
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

func addHooks(logger *logrus.Logger, config *Config) {
	if config.Hooks.Sentry != nil {
		sh, err := sentryHook(config.Hooks.Sentry)
		if err == nil {
			logger.AddHook(sh)
		} else {
			logger.WithError(err).Debug("can't add hook sentry")
		}
	}
}
