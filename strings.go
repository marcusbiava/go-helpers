package gohelpers

import (
	"fmt"
	"strings"
)

func sliceWithoutBracket[T any](slice []T) string {
	temp := fmt.Sprintf("%v", slice)
	split := strings.Split(temp[1:len(temp)-1], " ")
	join := strings.Join(split, ",")
	return fmt.Sprintf("%v", join)
}
