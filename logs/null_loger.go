package logs

type NullLog struct {
	LogWriter
}

func NewNullLog() *NullLog {
	return &NullLog{}
}

func (*NullLog) Debug(args ...interface{})   {}
func (*NullLog) Info(args ...interface{})    {}
func (*NullLog) Warning(args ...interface{}) {}
func (*NullLog) Error(args ...interface{})   {}
func (*NullLog) Fatal(args ...interface{})   {}
