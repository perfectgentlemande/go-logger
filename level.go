package logger

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
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
	// TraceLevel       // trace
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
	v, ok := levelNumber[strings.Replace(string(value), `"`, "", -1)]
	if !ok {
		return fmt.Errorf("wrong value of level: %s", string(value))
	}

	*level = v
	return nil
}

func (level *Level) UnmarshalYAML(value *yaml.Node) error {
	v, ok := levelNumber[value.Value]
	if !ok {
		return fmt.Errorf("wrong value of level: %s", value.Value)
	}

	level = &v
	return nil
}
