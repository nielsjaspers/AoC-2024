package day5

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func Day5() int {
    in, err := getFileInput("internal/input/day5.input")
    if err != nil {
        log.Fatalf("Error while getting file input %v", err)
    }
    rules, updates := parseInput(in)
    result := findMiddlePagesSum(rules, updates)
    return result
}

func getFileInput(path string) ([]string, error) {
    openedFile, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("Error while opening file: %v", err)
    }
    defer openedFile.Close()

    fileScanner := bufio.NewScanner(openedFile)
    fileScanner.Split(bufio.ScanLines)

    var fullLine []string

    for fileScanner.Scan() {
        fullLine = append(fullLine, fileScanner.Text())
    }

    return fullLine, nil
}

func parseInput(input []string) (map[int][]int, [][]int) {
    rules := make(map[int][]int)
    var updates [][]int
    isUpdateSection := false

    for _, line := range input {
        line = strings.TrimSpace(line)
        if line == "" {
            isUpdateSection = true
            continue
        }

        if isUpdateSection {
            parts := strings.Split(line, ",")
            var update []int
            for _, part := range parts {
                page, _ := strconv.Atoi(part)
                update = append(update, page)
            }
            updates = append(updates, update)
        } else {
            parts := strings.Split(line, "|")
            if len(parts) != 2 {
                continue
            }
            left, _ := strconv.Atoi(parts[0])
            right, _ := strconv.Atoi(parts[1])
            rules[left] = append(rules[left], right)
        }
    }

    return rules, updates
}

func isValidUpdate(update []int, rules map[int][]int) bool {
    position := make(map[int]int)
    for i, page := range update {
        position[page] = i
    }

    for page1, afterPages := range rules {
        if pos1, ok := position[page1]; ok {
            for _, page2 := range afterPages {
                if pos2, ok := position[page2]; ok {
                    if pos1 > pos2 {
                        return false
                    }
                }
            }
        }
    }

    return true
}

func findMiddlePage(update []int) int {
    return update[len(update)/2]
}

func findMiddlePagesSum(rules map[int][]int, updates [][]int) int {
    var sum int
    for _, update := range updates {
        if isValidUpdate(update, rules) {
            middlePage := findMiddlePage(update)
            sum += middlePage
        }
    }
    return sum
}
