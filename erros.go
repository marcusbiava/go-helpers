package gohelpers

import "log"

var logFatalf func(format string, v ...any) = log.Fatalf

func IfAnErrorOccursCallsLogFatal(message string, err error) {
	if err != nil {
		logFatalf("ERROR: %s \n%v", message, err)
	}
}
