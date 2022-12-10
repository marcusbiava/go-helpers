package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string
	Ids     []int
}

func main() {
	pigeon := Bird{
		Species: "Pigeon",
		Ids:     []int{1, 2, 3},
	}

	pigen2 := Bird{
		Species: "sdfafa",
		Ids:     []int{4, 5, 6},
	}

	// Now we pass a slice of two pigeons
	birds := []Bird{pigeon, pigen2}
	data, _ := json.Marshal(birds)
	fmt.Println(string(data))

}
