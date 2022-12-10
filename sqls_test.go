package gohelpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlIn(t *testing.T) {
	expected := `prop in (1,2,3)
or prop in (4,5,6)
or prop in (7,8,9)
or prop in (10)
`
	actual := SqlIn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "prop", 3)
	assert.Equal(t, expected, actual)

	fmt.Println(actual)
}
