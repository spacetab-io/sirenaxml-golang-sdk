package logs

import (
	"github.com/microparts/logs-go"
)

var Log *logs.Logger

func Init(logger *logs.Logger) (err error) {
	Log = logger
	return
}
