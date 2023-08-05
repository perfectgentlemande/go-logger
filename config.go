package logger

type Config struct {
	Level     Level  `yaml:"level" json:"level"`         // enum (panic|fatal|error|warning|info|debug|trace)
	Formatter string `yaml:"formatter" json:"formatter"` // enum (json|text)
	Output    string `yaml:"output" json:"output"`       // enum (stdout|stderr|path/to/file)
}
