package logger

type zerologWrapper struct{}

func newZerolog(config *Config) Logger {
	return &zerologWrapper{}
}

func (zw *zerologWrapper) Debug(args ...interface{})   {}
func (zw *zerologWrapper) Info(args ...interface{})    {}
func (zw *zerologWrapper) Warning(args ...interface{}) {}
func (zw *zerologWrapper) Error(args ...interface{})   {}
func (zw *zerologWrapper) Fatal(args ...interface{})   {}
