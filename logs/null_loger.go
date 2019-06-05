package logs

type NullLog struct {
	LogWriter
}

func NewNullLog() *NullLog {
	return &NullLog{}
}

func (*NullLog) Debug(args ...interface{})    {}
func (*NullLog) Debugf(args ...interface{})   {}
func (*NullLog) Info(args ...interface{})     {}
func (*NullLog) Infof(args ...interface{})    {}
func (*NullLog) Warning(args ...interface{})  {}
func (*NullLog) Warningf(args ...interface{}) {}
func (*NullLog) Error(args ...interface{})    {}
func (*NullLog) Errorf(args ...interface{})   {}
func (*NullLog) Fatal(args ...interface{})    { panic(args) }
func (*NullLog) Fatalf(args ...interface{})   { panic(args) }
