package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func main() {
	if err := runner.Run(problem1, problem2); err != nil {
		log.Fatal(err)
	}
}

func problem1(lines []string) (int, error) {
	rules, updates, err := readLines(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, update := range updates {
		if isCorrectOrder(rules, update) {
			count += middle(update)
		}
	}
	return count, nil
}

func problem2(lines []string) (int, error) {
	rules, updates, err := readLines(lines)
	if err != nil {
		return 0, nil
	}
	var count int
	for _, update := range updates {
		if !isCorrectOrder(rules, update) {
			count += reorder(rules, update)
		}
	}
	return count, nil
}

func reorder(rules [][]int, update []int) int {
	for i := range update {
		for j := i + 1; j < len(update); j++ {
			for _, r := range rules {
				if update[j] == r[0] && update[i] == r[1] {
					swap := update[i]
					update[i] = update[j]
					update[j] = swap
				}
			}
		}
	}
	return middle(update)
}

func readLines(lines []string) ([][]int, [][]int, error) {
	var lines1, lines2 []string
	for _, line := range lines {
		if strings.Contains(line, "|") {
			lines1 = append(lines1, line)
		}
		if strings.Contains(line, ",") {
			lines2 = append(lines2, line)
		}
	}
	rules, err := readRules(lines1)
	if err != nil {
		return nil, nil, err
	}
	updates, err := readUpdates(lines2)
	if err != nil {
		return nil, nil, err
	}
	return rules, updates, nil
}

func isCorrectOrder(rules [][]int, update []int) bool {
	for i := range update {
		for j := i + 1; j < len(update); j++ {
			for _, r := range rules {
				if update[j] == r[0] && update[i] == r[1] {
					return false
				}
			}
		}
	}
	return true
}

func readRules(input []string) ([][]int, error) {
	rules := [][]int{}
	for _, in := range input {
		r := strings.SplitN(in, "|", 2)
		i, err := strconv.Atoi(r[0])
		if err != nil {
			return nil, err
		}
		j, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, err
		}
		rules = append(rules, []int{i, j})
	}
	return rules, nil
}

func readUpdates(input []string) ([][]int, error) {
	var output [][]int
	for _, in := range input {
		pages := strings.Split(in, ",")
		var row []int
		for _, p := range pages {
			i, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		output = append(output, row)
	}
	return output, nil
}

func middle(row []int) int {
	return row[len(row)/2]
}
