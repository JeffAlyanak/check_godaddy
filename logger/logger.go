package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Logger struct bundles log.Logger and filename
type Logger struct {
	filename string
	*log.Logger
}

var lg *Logger
var once sync.Once

// Get Returns a Logger
func Get() *Logger { // TODO: Add user configurable log directory
	once.Do(func() {
		lg = makeLogger("check_godaddy.log")
	})
	return lg
}

func makeLogger(fname string) *Logger {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Could not open log file at: ", fname)
		os.Exit(3)
	}

	return &Logger{
		filename: fname,
		Logger:   log.New(file, "", log.Ldate|log.Ltime),
	}
}
