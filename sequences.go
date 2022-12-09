package gohelpers

import "sort"

func Sequence(slice []int) [][]int {
	var rangeSlice [][]int

	sort.Ints(slice)

	size := len(slice) - 1
	first := slice[0]

	for i, actual := range slice {
		next := slice[min(size, i+1)]
		sequential := next - actual

		if i == 0 && sequential == 2 {
			rangeSlice = addToRange(rangeSlice, actual, actual)
			first = next
			continue
		}

		if sequential == 0 || sequential == 2 {
			rangeSlice = addToRange(rangeSlice, first, actual)
			first = next
		}
	}
	return rangeSlice
}

func addToRange(rangeSlice [][]int, first int, second int) [][]int {
	var add []int
	if first == second {
		add = []int{first}
	} else {
		add = []int{first, second}
	}
	return append(rangeSlice, add)
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
