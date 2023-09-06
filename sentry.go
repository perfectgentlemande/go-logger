package logger

import (
	"github.com/sirupsen/logrus"
)

func logLevels(config *Sentry) []logrus.Level {
	levels := make([]logrus.Level, 0, config.Level+1)
	for i := PanicLevel; i <= config.Level; i++ {
		levels = append(levels, logrusLevelMap[i])
	}
	return levels
}
