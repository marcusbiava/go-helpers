package gohelpers

import (
	"os"
	"reflect"
	"testing"
)

func Test_files(t *testing.T) {
	fileName := "test_files/file1"
	os.Remove(fileName)

	data := []string{"1", "2", "3", "4"}

	WriteLines(data, fileName)

	newData := ReadLines(fileName)

	ok := reflect.DeepEqual(data, newData)

	if !ok {
		t.Errorf("Error in Test_files func")
	}

	os.Remove(fileName)
}
