package day1

import (
	"log"
	"slices"
)

func Day1p2() int {
	leftCol, rightCol, err := getFileInput("internal/input/day1_input.txt")
	if err != nil {
		log.Printf("Error while getting file input: %v", err)
	}

    slices.Sort(rightCol)

	result := listsToMaps(leftCol, rightCol)
	return result
}

func listsToMaps(leftCol []int, rightCol []int) int {
	m := make(map[int]int)

	var lastIndex int = 0

    	for _, left := range leftCol {
        	for i := lastIndex; i < len(rightCol); i++ {
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
