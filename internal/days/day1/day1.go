package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1() int {
    leftCol, rightCol, err := getFileInput("internal/input/day1_input.txt")
    if err != nil {
        log.Printf("Error while getting file input: %v", err)
    }

    distance := getDistance(leftCol, rightCol)

    return distance
}

func getFileInput(path string) ([]int, []int, error) {
	openedFile, err := os.Open(path)
	defer openedFile.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("Error while opening file: %v", err)
	}

	fileScanner := bufio.NewScanner(openedFile)
	fileScanner.Split(bufio.ScanLines)

	var fullLine []string

	for fileScanner.Scan() {
		fullLine = append(fullLine, fileScanner.Text())
	}

	leftCol, rightCol := splitStringToInt(fullLine)
    if len(leftCol) != len(rightCol) {
        log.Fatalln("Columns are not the same length. Something went wrong...")
    }

	return leftCol, rightCol, nil
}

func splitStringToInt(input []string) ([]int, []int) {
	var leftColInt, rightColInt []int

	for _, v := range input {
		splitLine := strings.Split(v, "   ")

		leftInt, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatalf("Error parsing string to int: %v", err)
		}
		leftColInt = append(leftColInt, leftInt)

		rightInt, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatalf("Error parsing string to int: %v", err)
		}
		rightColInt = append(rightColInt, rightInt)

	}

	return leftColInt, rightColInt
}

func getDistance(leftCol []int, rightCol []int) int {
	var distance int = 0

	// Sort arrays
	slices.Sort(leftCol)
	slices.Sort(rightCol)

    for i := range leftCol {
        if leftCol[i] < rightCol[i] {
            distance += rightCol[i] - leftCol[i]
        } else {
            distance += leftCol[i] - rightCol[i]
        }
    }

    return distance
}
