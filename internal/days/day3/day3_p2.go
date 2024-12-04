package day3

import (
	"log"
	"regexp"
	"strconv"
)

func Day3_p2() int {
	in, err := getFileInput("internal/input/day3.input")
	if err != nil {
		log.Fatalf("Error while getting file input: %v", err)
	}
	total := regexP2(in)
	return total
}

func regexP2(list []string) int {
	patternMul := `mul\((\d+),(\d+)\)`
	patternDo := `do\(\)`
	patternDont := `don't\(\)`

	reMul := regexp.MustCompile(patternMul)
	reDo := regexp.MustCompile(patternDo)
	reDont := regexp.MustCompile(patternDont)

	// Initialize state and result
	enabled := true // At the start, mul instructions are enabled
	result := 0

	for _, str := range list {
		// Split the line into separate instructions using regex
		instructionRegex := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
		instructions := instructionRegex.FindAllString(str, -1)

		for _, instr := range instructions {
			// Check for "do()" instruction
			if reDo.MatchString(instr) {
				enabled = true
				continue
			}

			// Check for "don't()" instruction
			if reDont.MatchString(instr) {
				enabled = false
				continue
			}

			if reMul.MatchString(instr) {
				matches := reMul.FindStringSubmatch(instr)
				x, _ := strconv.Atoi(matches[1])
				y, _ := strconv.Atoi(matches[2])

				if enabled {
					result += x * y
				}
			}
		}
	}

	return result
}
