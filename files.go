package gohelpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	IfAnErrorOccursCallsLogFatal("readLines", err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func WriteLines(lines []string, path string) {
	file, err := os.Create(path)
	IfAnErrorOccursCallsLogFatal("writeLines create", err)
	defer func(file *os.File) {
		err := file.Close()
		IfAnErrorOccursCallsLogFatal("close", err)
	}(file)

	w := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := fmt.Fprintln(w, line)
		IfAnErrorOccursCallsLogFatal("Fprintln", err)
	}
	err = w.Flush()
	IfAnErrorOccursCallsLogFatal("Flush", err)
}

func WriteStringToEndOfFile(str, path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	IfAnErrorOccursCallsLogFatal("Could not open "+path, err)
	defer func(file *os.File) {
		err := file.Close()
		IfAnErrorOccursCallsLogFatal("close", err)
	}(file)

	_, err2 := file.WriteString(str)
	IfAnErrorOccursCallsLogFatal("Could not write text to "+path, err2)
}
