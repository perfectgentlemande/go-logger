package logger

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Level     Level     `yaml:"level" json:"level"`         // enum (panic|fatal|error|warning|info|debug|trace)
	Formatter Formatter `yaml:"formatter" json:"formatter"` // enum (json|text)
	Output    string    `yaml:"output" json:"output"`       // enum (stdout|stderr|path/to/file)
}

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

	*level = v
	return nil
}

type Formatter string

const (
	FormatterJSON Formatter = "json"
	FormatterText Formatter = "text"
)

func (f *Formatter) isValid() bool {
	_, ok := map[Formatter]struct{}{
		FormatterJSON: {},
		FormatterText: {},
	}[*f]

	return ok
}

func (f *Formatter) UnmarshalJSON(value []byte) error {
	v := Formatter(strings.Replace(string(value), `"`, "", -1))
	ok := v.isValid()

	if !ok {
		return fmt.Errorf("wrong value of formatter: %s", v)
	}

	*f = v
	return nil
}

func (f *Formatter) UnmarshalYAML(value *yaml.Node) error {
	v := Formatter(value.Value)
	ok := v.isValid()

	if !ok {
		return fmt.Errorf("wrong value of level: %s", value.Value)
	}

	*f = v
	return nil
}
