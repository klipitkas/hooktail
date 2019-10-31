package logging

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// Log is the common shared logrus instance.
var Log = log.New()

// The default file for logging.
var logFile = "hooktail.log"

func init() {
	// Open the log file and try writing to it.
	logFile, err := os.OpenFile(logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)

	if err != nil {
		Log.Fatalf("write to log file %s: %v", logFile.Name(), err)
	}

	// Write logs to both os.Stdout and logfile
	mw := io.MultiWriter(os.Stdout, logFile)

	// Write to both outputs.
	Log.SetOutput(mw)

	// Set the default log level.
	Log.SetLevel(log.InfoLevel)

	// Set the default formatter options.
	Log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
}
