package gohelpers

import (
	"testing"
)

func TestContainsTrue(t *testing.T) {
	slice := []string{"1", "2"}
	ok := Contains(slice, "1")

	if ok == false {
		t.Errorf("Error")
	}
}

func TestContainsFalse(t *testing.T) {
	slice := []string{"1", "2"}
	ok := Contains(slice, "3")

	if ok == true {
		t.Errorf("Error")
	}
}

func TestRemove(t *testing.T) {
	slice := []string{"1", "2"}
	newSlice := Remove(slice, "2")

	if !Contains(newSlice, "1") {
		t.Errorf("Error")
	}

	if Contains(newSlice, "2") {
		t.Errorf("Error")
	}
}
