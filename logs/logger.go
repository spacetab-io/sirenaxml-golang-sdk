package logs

var Logger LogWriter

type LogWriter interface {
	Debug(msg interface{})
	Info(msg interface{})
	Warning(msg interface{})
	Error(msg interface{})
	Fatal(msg interface{})
}
