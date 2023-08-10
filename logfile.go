package logfile

import (
	"log"
	"os"
)

var (
	logFile  *os.File
	logError error
)

func LogFile(logFilePath string) *os.File {
	logFile, logError = os.OpenFile(
		logFilePath,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0664,
	)
	if logError != nil {
		os.Exit(1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	return logFile
}
