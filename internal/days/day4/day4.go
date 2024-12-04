package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day4() int {
    in, err := getFileInput("internal/input/day4.input")
    if err != nil {
        log.Fatalf("Error while getting file input %v", err)
    }
    result := countOccurrences(in)
    return result
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

// Directions for traversal
var directions = [8][2]int{
	{0, 1},  // Right
	{0, -1}, // Left
	{1, 0},  // Down
	{-1, 0}, // Up
	{1, 1},  // Down-Right Diagonal
	{1, -1}, // Down-Left Diagonal
	{-1, 1}, // Up-Right Diagonal
	{-1, -1}, // Up-Left Diagonal
}

func countOccurrences(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	word := "XMAS"
	wordLen := len(word)
	count := 0

	// Function to check if the word matches in a specific direction
	checkWord := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLen; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != word[i] {
				return false
			}
		}
		return true
	}

	// Traverse the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Check each direction
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkWord(i, j, dx, dy) {
					count++
				}
			}
		}
	}

	return count
}
