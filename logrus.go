package logger

type logrusWrapper struct{}

func newLogrus(config *Config) Logger {
	return &logrusWrapper{}
}

func (lw *logrusWrapper) Debug(args ...interface{})   {}
func (lw *logrusWrapper) Info(args ...interface{})    {}
func (lw *logrusWrapper) Warning(args ...interface{}) {}
func (lw *logrusWrapper) Error(args ...interface{})   {}
func (lw *logrusWrapper) Fatal(args ...interface{})   {}
