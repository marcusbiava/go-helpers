package gohelpers

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Remove[T comparable](s []T, e T) []T {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func ChunkSlice[T any](input []T, size int) [][]T {
	var chunks [][]T

	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[i:end])
	}
	return chunks
}

func RemoveDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	var list []T
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Intersection[T comparable](slices ...[]T) []T {
	possibleIntersections := map[T]int{}
	for i, slice := range slices {
		for _, el := range slice {
			if i == 0 {
				possibleIntersections[el] = 0
			} else if _, elementExists := possibleIntersections[el]; elementExists {
				possibleIntersections[el] = i
			}
		}
	}

	intersected := make([]T, 0)
	for _, el := range slices[0] {
		if lastVisitorIndex, exists := possibleIntersections[el]; exists && lastVisitorIndex == len(slices)-1 {
			intersected = append(intersected, el)
			delete(possibleIntersections, el)
		}
	}

	return intersected
}

func Difference[T comparable](slices ...[]T) []T {
	possibleDifferences := map[T]int{}
	nonDifferentElements := map[T]int{}

	for i, slice := range slices {
		for _, el := range slice {
			if lastVisitorIndex, elementExists := possibleDifferences[el]; elementExists && lastVisitorIndex != i {
				nonDifferentElements[el] = i
			} else if !elementExists {
				possibleDifferences[el] = i
			}
		}
	}

	differentElements := make([]T, 0)

	for _, slice := range slices {
		for _, el := range slice {
			if _, exists := nonDifferentElements[el]; !exists {
				differentElements = append(differentElements, el)
			}
		}
	}

	return differentElements
}

func Map[T any, R any](slice []T, mapper func(value T, index int, slice []T) R) (mapped []R) {
	for i, el := range slice {
		mapped = append(mapped, mapper(el, i, slice))
	}
	return mapped
}
