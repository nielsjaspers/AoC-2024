package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day3() int {
    in, err := getFileInput("internal/input/day3.mock")
    if err != nil {
        log.Fatalf("Error getting file input: %v", err)
    }
    total := regex(in)
    return total
}

func getFileInput(path string) ([]string, error) {
	openedFile, err := os.Open(path)
	defer openedFile.Close()
	if err != nil {
		return nil, fmt.Errorf("Error while opening file: %v", err)
	}

	fileScanner := bufio.NewScanner(openedFile)
	fileScanner.Split(bufio.ScanLines)

	var fullLine []string

	for fileScanner.Scan() {
		fullLine = append(fullLine, fileScanner.Text())
	}

    return fullLine, nil
}

func regex(list []string) int {
	pattern := `mul\((\d+),(\d+)\)`

	re := regexp.MustCompile(pattern)

	var results int

	for _, str := range list {
		matches := re.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1]) 
			y, _ := strconv.Atoi(match[2]) 
			results += x * y
		}
	}

	return results
}
