package logger

import "log"

func Info(msg string, arg...interface{}) {
	log.Printf(msg, arg...)
}
