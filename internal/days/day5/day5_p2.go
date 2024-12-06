package day5

import "log"

func Day5_p2() int {
    in, err := getFileInput("internal/input/day5.input")
    if err != nil {
        log.Fatalf("Error while getting file input: %v\n", err)
    }
    rules, updates := parseInput(in)
    incorrectUpdates := findIncorrectUpdates(rules, updates)
    result := findMiddlePagesSumForIncorrectUpdates(rules, incorrectUpdates)
    return result
}

func findIncorrectUpdates(rules map[int][]int, updates [][]int) [][]int {
    var incorrectUpdates [][]int
    for _, update := range updates {
        if !isValidUpdate(update, rules) {
            incorrectUpdates = append(incorrectUpdates, update)
        }
    }
    return incorrectUpdates
}

func sortUpdate(update []int, rules map[int][]int) []int {
    graph := make(map[int][]int)
    inDegree := make(map[int]int)
    for _, page := range update {
        graph[page] = []int{}
        inDegree[page] = 0
    }

    for page1, afterPages := range rules {
        for _, page2 := range afterPages {
            if contains(update, page1) && contains(update, page2) {
                graph[page1] = append(graph[page1], page2)
                inDegree[page2]++
            }
        }
    }

    var sortedUpdate []int
    queue := []int{}
    for page, degree := range inDegree {
        if degree == 0 {
            queue = append(queue, page)
        }
    }

    for len(queue) > 0 {
        page := queue[0]
        queue = queue[1:]
        sortedUpdate = append(sortedUpdate, page)
        for _, neighbor := range graph[page] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    return sortedUpdate
}

func contains(slice []int, value int) bool {
    for _, v := range slice {
        if v == value {
            return true
        }
    }
    return false
}

func findMiddlePagesSumForIncorrectUpdates(rules map[int][]int, updates [][]int) int {
    var sum int
    for _, update := range updates {
        sortedUpdate := sortUpdate(update, rules)
        middlePage := findMiddlePage(sortedUpdate)
        sum += middlePage
    }
    return sum
}
