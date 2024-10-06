package logging

import (
	"log"
	"os"
)

var (
	warn  *log.Logger
	info  *log.Logger
	error *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warn = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(v ...any) {
	info.Println(v...)
}
func Error(v ...any) {
	error.Println(v...)
}
func Warn(v ...any) {
	warn.Println(v...)
}
