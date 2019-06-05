package logs

var Logger LogWriter

type LogWriter interface {
	Debug(args ...interface{})
	Debugf(args ...interface{})
	Info(args ...interface{})
	Infof(args ...interface{})
	Warning(args ...interface{})
	Warningf(args ...interface{})
	Error(args ...interface{})
	Errorf(args ...interface{})
	Fatal(args ...interface{})
	Fatalf(args ...interface{})
}
