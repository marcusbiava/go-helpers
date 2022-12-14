package gohelpers

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestContainsTrue(t *testing.T) {
	sliceStr := []string{"1", "2"}

	assert.Equal(t, true, Contains(sliceStr, "1"))
	assert.Equal(t, false, Contains(sliceStr, "3"))

	sliceInt := []int{1, 2}
	assert.Equal(t, true, Contains(sliceInt, 1))
	assert.Equal(t, false, Contains(sliceInt, 3))
}

func TestRemove(t *testing.T) {
	assert.Equal(t, []string{"1", "3"}, Remove([]string{"1", "2", "3"}, "2"))
	assert.Equal(t, []int{1, 3}, Remove([]int{1, 2, 3}, 2))
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
