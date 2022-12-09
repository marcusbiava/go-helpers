package gohelpers

import (
	"fmt"
	"strings"
)

func SqlIn[T any](s []T, prop string, chunkSize int) string {
	var sb strings.Builder
	var chunks = ChunkSlice(s, chunkSize)

	for i, v := range chunks {
		if i != 0 {
			sb.WriteString("or ")
		}
		sb.WriteString(prop + " in ")

		a := fmt.Sprintf("%v", v)
		a = strings.Replace(a, " ", ",", -1)
		a = strings.Replace(a, "[", "(", -1)
		a = strings.Replace(a, "]", ")", -1)

		sb.WriteString(a)
		sb.WriteString("\n")

	}
	return sb.String()
}
