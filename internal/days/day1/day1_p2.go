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

	slices.Sort(leftCol)
	slices.Sort(rightCol)

	result := listsToMaps(leftCol, rightCol)
	return result
}

func listsToMaps(leftCol []int, rightCol []int) int {
	m := make(map[int]int)
	originalValues := make(map[int]int)

	var lastIndex int = 0
	var prevVal int = -1

    var index int = 0

	for _, left := range leftCol {
		for i := lastIndex; i < len(rightCol); i++ {
            index++
			//fmt.Printf("l: %v\t\tr: %v\n", left, rightCol[i])

			if rightCol[i] == left {
				m[left]++
			}

			if rightCol[i] > left {
				lastIndex = i
				break
			}
		}

		// check for Duplicate
		if prevVal == left {
			m[left] += originalValues[left]
		} else {
			originalValues[left] = m[left]
		}

		prevVal = left

	}

	total := 0
	for key, val := range m {
		total += key * val
	}
    fmt.Printf("Total times looped: %v\n", index)
	return total
}
