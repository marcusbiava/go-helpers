package gohelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SliceWithoutBracket_Int(t *testing.T) {
	numerals := []int{1, 3, 4}

	actualResult := SliceWithoutBracket(numerals)
	expected := "1,3,4"
	assert.Equal(t, expected, actualResult)
}

func Test_SliceWithoutBracket_String(t *testing.T) {
	numerals := []string{"1", "3", "4"}

	actualResult := SliceWithoutBracket(numerals)
	expected := "1,3,4"
	assert.Equal(t, expected, actualResult)
}
