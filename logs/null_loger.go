package logs

type NullLog struct {
	LogWriter
}

func NewNullLog() *NullLog {
	return &NullLog{}
}

func (*NullLog) Debug(msg interface{})   {}
func (*NullLog) Notice(msg interface{})  {}
func (*NullLog) Info(msg interface{})    {}
func (*NullLog) Warning(msg interface{}) {}
func (*NullLog) Error(msg interface{})   {}
func (*NullLog) Fatal(msg interface{})   {}
