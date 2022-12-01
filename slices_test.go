package gohelpers

import "testing"

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
