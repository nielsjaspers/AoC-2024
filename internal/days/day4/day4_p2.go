package day4

import "log"

func Day4_p2() int {
	in, err := getFileInput("internal/input/day4.input")
	if err != nil {
		log.Fatalf("Error while getting file input: %v\n", err)
	}
	result := countOccurrencesP2(in)
	return result
}

func countOccurrencesP2(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Iterate through the inner cells (skip outer edges)
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			// Check if the center is 'A'
			if grid[row][col] == 'A' {
				// Check top-left to bottom-right diagonal
				tl_br := isMAS(grid, row-1, col-1, row, col, row+1, col+1)
				// Check top-right to bottom-left diagonal
				tr_bl := isMAS(grid, row-1, col+1, row, col, row+1, col-1)

				// If both diagonals form MAS/SAM, count it
				if tl_br && tr_bl {
					count++
				}
			}
		}
	}

	return count
}

// Check if a diagonal forms "MAS" or "SAM"
func isMAS(grid []string, x1, y1, x2, y2, x3, y3 int) bool {
	if grid[x1][y1] == 'M' && grid[x2][y2] == 'A' && grid[x3][y3] == 'S' {
		return true
	}
	if grid[x1][y1] == 'S' && grid[x2][y2] == 'A' && grid[x3][y3] == 'M' {
		return true
	}
	return false
}
