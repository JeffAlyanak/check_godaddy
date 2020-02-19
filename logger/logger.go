package logger

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var lg *logger
var once sync.Once

func Get() *logger {
	once.Do(func() {
		lg = makeLogger("godaddy-check.log")
	})
	return lg
}

func makeLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "", log.Ldate|log.Ltime),
	}
}
