package day1

import (
	"fmt"
	"log"
	"slices"
)

func Day1p2() int {
	leftCol, rightCol, err := getFileInput("internal/input/day1_input.txt")
	if err != nil {
		log.Printf("Error while getting file input: %v", err)
	}

    slices.Sort(rightCol)

    fmt.Printf("leftCol: %v\nrightCol: %v", leftCol, rightCol)

	result := listsToMaps(leftCol, rightCol)
	return result
}

func listsToMaps(leftCol []int, rightCol []int) int {
	// for loop
	// select first item from leftCol, loop through rightCol until rightCol > leftCol
	// select second item form leftCol, start from previous endpoint and loop until rightCol > leftCol
	// repeat until all items in leftCol are mapped
	// loop once through all values gathered and do key * value
	// return result

	m := make(map[int]int)

	var lastIndex int = 0

    for _, left := range leftCol {
        for i := lastIndex; i < len(rightCol); i++ {
            fmt.Printf("current index: %v\t\tlast index: %v\n", i, lastIndex)

            fmt.Printf("l: %v\t\tr: %v\n", left, rightCol[i])

            if rightCol[i] == left {
                m[left]++
            }
        }
    }

	var total int = 0

	for key, val := range m {
		total += key * val
	}

	return total
}
