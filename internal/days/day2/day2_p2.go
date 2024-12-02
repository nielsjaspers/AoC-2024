package day2

import (
	"fmt"
)

func Day2_p2() int {
	file, err := getFileInput("internal/input/day2_input.txt")
	if err != nil {
		fmt.Printf("Error getting file input: %v", err)
	}

	var amountSafe int

	for _, line := range file {
		if isValidLineP2(line) {
			amountSafe++
		}
	}
	return amountSafe
}

func isValidLineP2(line []int) bool {
	if len(line) < 2 {
		return false
	}

	// Check if line is valid without removing any level
	if isStrictlyValid(line) {
		return true
	}

	// Try removing each level and check if the line becomes valid
	for i := 0; i < len(line); i++ {
		newLine := append([]int{}, line[:i]...)
		newLine = append(newLine, line[i+1:]...)
		if isStrictlyValid(newLine) {
			return true
		}
	}

	return false
}

func isStrictlyValid(line []int) bool {
	if len(line) < 2 {
		return false
	}

	isIncreasing := line[1] > line[0]

	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]

		if diff < -3 || diff > 3 {
			return false
		}

		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			return false
		}
	}

	return true
}

