package logger

import (
	"os"
	"sync"

	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/sirena-agent-go/config"

	logging "github.com/op/go-logging"
)

// Logger is a server logger
type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Critical(args ...interface{})
	Criticalf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Notice(args ...interface{})
	Noticef(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

// LoggerID ia a logger ID
const LoggerID = "Sirena SDK Logger"

var (
	// Logger settings
	logger           = logging.MustGetLogger(LoggerID)
	logConsoleFormat = logging.MustStringFormatter(
		`%{color}%{time:2006/01/02 15:04:05} (%{shortfile}) >> %{message} %{color:reset}`,
	)
)

var once sync.Once

// Get returns the logger
func Get() Logger {
	once.Do(func() {
		// Prepare logger
		logConsoleBackend := logging.NewLogBackend(os.Stderr, "", 0)
		logConsolePrettyBackend := logging.NewBackendFormatter(logConsoleBackend, logConsoleFormat)
		logging.SetBackend(logConsolePrettyBackend)

		conf := config.Get()

		// Set proper log level based on config
		switch conf.LogLevel {
		case "debug":
			logging.SetLevel(logging.DEBUG, LoggerID)
		case "info":
			logging.SetLevel(logging.INFO, LoggerID)
		case "warn":
			logging.SetLevel(logging.WARNING, LoggerID)
		case "err":
			logging.SetLevel(logging.ERROR, LoggerID)
		default:
			logging.SetLevel(logging.DEBUG, LoggerID) // log everything by default
		}
	})
	return logger
}
