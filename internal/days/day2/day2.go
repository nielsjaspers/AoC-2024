package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() int {
	file, err := getFileInput("internal/input/day2_input.txt")
	if err != nil {
		fmt.Printf("Error getting file input: %v", err)
	}

	var amountSafe int

	for _, line := range file {
		if isValidLine(line) {
			amountSafe++
		}
	}
	return amountSafe
}

func getFileInput(path string) ([][]int, error) {
	openedFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error while opening file: %v", err)
	}
	defer openedFile.Close()

	fileScanner := bufio.NewScanner(openedFile)
	fileScanner.Split(bufio.ScanLines)

	var result [][]int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums, err := parseLineToInts(line)
		if err != nil {
			return nil, fmt.Errorf("Error parsing line to integers: %v", err)
		}
		result = append(result, nums)
	}

	if err := fileScanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return result, nil
}

func parseLineToInts(line string) ([]int, error) {
	parts := strings.Fields(line)
	var nums []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("Invalid number: %v", part)
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func isValidLine(line []int) bool {
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
