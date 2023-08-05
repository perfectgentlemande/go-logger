package logger

import (
	"encoding/json"
	"reflect"
	"testing"
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
			TestDocument: []byte(`{"level":""}`),
			Expected:     nil,
		},
		{
			TestDocument: []byte(`{}`),
			Expected:     nil,
		},
		{
			TestDocument: []byte(`{"level":""}`),
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
			Level *Level `json:"level"`
		}{}

		err := json.Unmarshal(cases[i].TestDocument, &gotDoc)
		if err != nil && !cases[i].ExpectedError {
			t.Error(err)
			continue
		}

		if !reflect.DeepEqual(cases[i].Expected, gotDoc.Level) {
			t.Errorf("cases: %v, expected: %v, got: %v", i, cases[i].Expected, gotDoc.Level)
			continue
		}
	}
}
