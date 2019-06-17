package logs

type NullLog struct {
	LogWriter
}

func NewNullLog() *NullLog {
	return &NullLog{}
}

func (*NullLog) Debug(args ...interface{})                   {}
func (*NullLog) Debugf(format string, args ...interface{})   {}
func (*NullLog) Info(args ...interface{})                    {}
func (*NullLog) Infof(format string, args ...interface{})    {}
func (*NullLog) Warning(args ...interface{})                 {}
func (*NullLog) Warningf(format string, args ...interface{}) {}
func (*NullLog) Error(args ...interface{})                   {}
func (*NullLog) Errorf(format string, args ...interface{})   {}
func (*NullLog) Fatal(args ...interface{})                   { panic(args) }
func (*NullLog) Fatalf(format string, args ...interface{})   { panic(args) }
func (*NullLog) Write(p []byte) (n int, err error)           { return 0, nil }
