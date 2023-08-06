package logger

import (
	"encoding/json"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func levelToPtr(level Level) *Level {
	return &level
}

func TestLevelUnmarshalJSON(t *testing.T) {
	type TestCase struct {
		TestDocument  []byte
		Expected      *Level
		ExpectedError bool
	}

	cases := []TestCase{
		{
			TestDocument:  []byte(`{"level":"panic"}`),
			Expected:      levelToPtr(PanicLevel),
			ExpectedError: false,
		},
		{
			TestDocument:  []byte(`{"level":"fatal"}`),
			Expected:      levelToPtr(FatalLevel),
			ExpectedError: false,
		},
		{
			TestDocument: []byte(`{"level":"error"}`),
			Expected:     levelToPtr(ErrorLevel),
		},
		{
			TestDocument: []byte(`{"level":"warning"}`),
			Expected:     levelToPtr(WarnLevel),
		},
		{
			TestDocument: []byte(`{"level":"info"}`),
			Expected:     levelToPtr(InfoLevel),
		},
		{
			TestDocument: []byte(`{"level":"debug"}`),
			Expected:     levelToPtr(DebugLevel),
		},
		{
			TestDocument:  []byte(`{"level":""}`),
			Expected:      nil,
			ExpectedError: true,
		},
		{
			TestDocument: []byte(`{}`),
			Expected:     nil,
		},
		{
			TestDocument:  []byte(`{"level":"asfasf"}`),
			Expected:      nil,
			ExpectedError: true,
		},
	}

	for i := range cases {
		gotDoc := struct {
			Level *Level `yaml:"level"`
		}{}

		err := json.Unmarshal(cases[i].TestDocument, &gotDoc)
		if err != nil && !cases[i].ExpectedError {
			t.Errorf("cases: %v unexpected error: %v", i, err)
			continue
		}

		if !reflect.DeepEqual(cases[i].Expected, gotDoc.Level) && !cases[i].ExpectedError {
			t.Errorf("cases: %v, expected: %v, got: %v", i, cases[i].Expected, gotDoc.Level)
			continue
		}
	}
}

func TestLevelUnmarshalYAML(t *testing.T) {
	type TestCase struct {
		TestDocument  []byte
		Expected      *Level
		ExpectedError bool
	}

	cases := []TestCase{
		{
			TestDocument:  []byte(`level: panic`),
			Expected:      levelToPtr(PanicLevel),
			ExpectedError: false,
		},
		{
			TestDocument:  []byte(`level: fatal`),
			Expected:      levelToPtr(FatalLevel),
			ExpectedError: false,
		},
		{
			TestDocument: []byte(`level: error`),
			Expected:     levelToPtr(ErrorLevel),
		},
		{
			TestDocument: []byte(`level: warning`),
			Expected:     levelToPtr(WarnLevel),
		},
		{
			TestDocument: []byte(`level: info`),
			Expected:     levelToPtr(InfoLevel),
		},
		{
			TestDocument: []byte(`level: debug`),
			Expected:     levelToPtr(DebugLevel),
		},
		{
			TestDocument:  []byte(`level:`),
			Expected:      nil,
			ExpectedError: true,
		},
		{
			TestDocument: []byte(``),
			Expected:     nil,
		},
		{
			TestDocument:  []byte(`level: asfasf`),
			Expected:      nil,
			ExpectedError: true,
		},
	}

	for i := range cases {
		gotDoc := struct {
			Level *Level `yaml:"level"`
		}{}

		err := yaml.Unmarshal(cases[i].TestDocument, &gotDoc)
		if err != nil && !cases[i].ExpectedError {
			t.Errorf("cases: %v unexpected error: %v", i, err)
			continue
		}

		if !reflect.DeepEqual(cases[i].Expected, gotDoc.Level) && !cases[i].ExpectedError {
			t.Errorf("cases: %v, expected: %v, got: %v", i, cases[i].Expected, gotDoc.Level)
			continue
		}
	}
}
