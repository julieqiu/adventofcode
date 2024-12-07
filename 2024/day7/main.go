package main

import (
	"fmt"
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
	rows, err := parseInput(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, row := range rows {
		if compute(row[0], row[1:], false) {
			count += row[0]
		}
	}
	return count, nil
}

func problem2(lines []string) (int, error) {
	rows, err := parseInput(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, row := range rows {
		if compute(row[0], row[1:], true) {
			count += row[0]
		}
	}
	return count, nil
}

func parseInput(lines []string) ([][]int, error) {
	var rows [][]int
	for _, line := range lines {
		nums := strings.Fields(line)
		var numbers []int
		for _, n := range nums {
			v, err := strconv.Atoi(strings.TrimSuffix(n, ":"))
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, v)
		}
		rows = append(rows, numbers)
	}
	return rows, nil
}

func compute(key int, numbers []int, withConcat bool) bool {
	queue := []int{numbers[0]}
	for i := 1; i < len(numbers); i++ {
		var newqueue []int
		for j := range queue {
			newqueue = append(newqueue, numbers[i]+queue[j])
			newqueue = append(newqueue, numbers[i]*queue[j])
			if withConcat {
				n := fmt.Sprintf("%d%d", queue[j], numbers[i])
				v, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				newqueue = append(newqueue, v)
			}
		}
		queue = newqueue
	}
	for _, r := range queue {
		if r == key {
			return true
		}
	}
	return false
}
