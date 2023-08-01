package logger

import (
	"fmt"
)

type Level uint32

const (
	emptyLevel Level = iota
	PanicLevel       // panic
	FatalLevel       // fatal
	ErrorLevel       // error
	WarnLevel        // warning
	InfoLevel        // info
	DebugLevel       // debug
	TraceLevel       // trace
)

var levelNumber = map[string]Level{
	"panic":   PanicLevel,
	"fatal":   FatalLevel,
	"error":   ErrorLevel,
	"warning": WarnLevel,
	"info":    InfoLevel,
	"debug":   DebugLevel,
	// "trace":   TraceLevel,
}

func (level *Level) UnmarshalJSON(value []byte) error {
	v, ok := levelNumber[string(value)]
	if !ok {
		return fmt.Errorf("wrong value of level: %s", string(value))
	}

	level = &v
	return nil
}
