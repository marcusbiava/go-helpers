package gohelpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	IfAnErrorOccursCallsLogFatal("readLines", err)
	defer file.Close()

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
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	w.Flush()
}
