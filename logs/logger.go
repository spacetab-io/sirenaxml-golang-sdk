package logs

import (
	"github.com/microparts/logs-go"
)

var Log *logs.Logger

func Init(logsCfg *logs.Config) (err error) {
	//logsCfg.Level = "debug"
	if Log, err = logs.NewLogger(logsCfg); err != nil {
		return err
	}
	return
}
