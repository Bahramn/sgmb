package logger

import (
	"github.com/Bahramn/sgmb/config"
	"log"
	"os"
)

func InitLogger(conf config.LoggerConfig) {
	// create logfile
	n := conf.Path + "/sgmb.log"
	f, err := os.OpenFile(n, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("unable to open log file with path: %s  %v", n, err)
	}
	log.SetOutput(f)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
