package gohelpers

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteStringToEndOfFile(t *testing.T) {
	fileName := "test_files/file1"

	data := []string{"1", "2", "3", "4"}

	WriteLines(data, fileName)

	newData := ReadLines(fileName)

	assert.Equal(t, 4, len(newData))

	WriteStringToEndOfFile("5", fileName)

	newData = ReadLines(fileName)

	assert.Equal(t, 5, len(newData))

	err := os.Remove(fileName)
	IfAnErrorOccursCallsLogFatal("remove", err)
}
