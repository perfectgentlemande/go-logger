package logger

type Config struct {
	Level     Level  `yaml:"level" json:"level" toml:"level"`             // enum (panic|fatal|error|warning|info|debug|trace)
	Formatter string `yaml:"formatter" json:"formatter" toml:"formatter"` // enum (json|text)
}
