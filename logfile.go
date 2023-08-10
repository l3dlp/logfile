package logfile

import (
	"fmt"
	"log"
	"os"
)

var (
	logFile  *os.File
	logError error
)

// The function opens a log file, sets it as the output for the log package, and returns the file.
func Use(logFilePath string) *os.File {
	logFile, logError = os.OpenFile(
		logFilePath,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0664,
	)
	if logError != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier de log:", logError)
		return nil
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	return logFile
}
