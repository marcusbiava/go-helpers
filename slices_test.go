package gohelpers

import (
	"github.com/stretchr/testify/assert"
	"strconv"
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

func TestChunkSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}, ChunkSlice(numbers, 2))
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}, ChunkSlice(numbers, 3))
}

func TestRemoveDuplicate(t *testing.T) {
	numbers := []int{1, 2, 1, 3, 5, 1, 7, 8, 3, 5}

	assert.Equal(t, []int{1, 2, 3, 5, 7, 8}, RemoveDuplicate(numbers))
}

func TestIntersection(t *testing.T) {
	numbers := []int{1, 2, 1, 3, 5, 1, 7, 8, 3, 5}
	numbers2 := []int{1, 2, 3, 99}

	assert.Equal(t, []int{1, 2, 3}, Intersection(numbers, numbers2))
}

func TestDifference(t *testing.T) {
	numbers := []int{1, 2, 1, 3, 5, 1, 7, 8, 3, 5}
	numbers2 := []int{1, 2, 3, 99}

	assert.Equal(t, []int{5, 7, 8, 5, 99}, Difference(numbers, numbers2))
}

func TestMap(t *testing.T) {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numeralsStr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	actualResult := Map(numerals, func(value int, _ int, _ []int) string {
		return strconv.Itoa(value)
	})
	assert.Equal(t, numeralsStr, actualResult)

	actualResult2 := Map(numeralsStr, func(value string, index int, slice []string) int {
		i, _ := strconv.Atoi(value)
		return i
	})

	assert.Equal(t, numerals, actualResult2)
}
