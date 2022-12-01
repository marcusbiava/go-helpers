package gohelpers

import (
	"errors"
	"log"
	"testing"
)

func TestIfAnErrorOccursCallsLogFatal(t *testing.T) {
	var ok bool

	logFatalf = func(format string, v ...any) {
		ok = true
	}

	IfAnErrorOccursCallsLogFatal("Test", errors.New("New errors"))

	if ok == false {
		t.Errorf("Error in AbstractFunction()")
	}

	logFatalf = log.Fatalf
}
