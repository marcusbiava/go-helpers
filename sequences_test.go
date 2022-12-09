package gohelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequencesCase1(t *testing.T) {
	numerals := []int{3, 1, 2, 0}
	actualResult := Sequence(numerals)
	expected := [][]int{
		{0, 3},
	}
	assert.Equal(t, expected, actualResult)
}

func TestSequencesCase2(t *testing.T) {
	numerals := []int{0, 1, 2, 9, 5, 6, 7, 8, 3}
	actualResult := Sequence(numerals)
	expected := [][]int{
		{0, 3},
		{5, 9},
	}
	assert.Equal(t, expected, actualResult)
}

func TestSequencesCase3(t *testing.T) {
	numerals := []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	actualResult := Sequence(numerals)
	expected := [][]int{
		{1},
		{3, 12},
	}
	assert.Equal(t, expected, actualResult)
}

func TestSequencesCase4(t *testing.T) {
	numerals := []int{1,
		3, 4, 5, 6, 7, 8,
		10,
		12, 13, 14, 15,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
		40, 41, 44, 45, 46}
	actualResult := Sequence(numerals)
	expected := [][]int{
		{1},
		{3, 8},
		{10},
		{12, 15},
		{17, 38},
		{40, 46},
	}
	assert.Equal(t, expected, actualResult)
}
