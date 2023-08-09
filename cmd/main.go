package main

import (
	"github.com/perfectgentlemande/go-logger"
)

func main() {
	log := logger.NewZerolog(&logger.Config{
		Level:     logger.DebugLevel,
		Formatter: logger.FormatterJSON,
		Output:    logger.OutputStdOut,
	}).WithFields(logger.Fields{
		"field_a": "a",
		"field_b": "b",
		"field_c": "c",
	})

	log.Info("Hello")
}
