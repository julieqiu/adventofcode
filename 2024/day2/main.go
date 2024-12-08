package main

import (
	"log"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func main() {
	if err := runner.Run(problem1, problem2); err != nil {
		log.Fatal(err)
	}
}

func problem1(lines []string) (int, error) {
	lists, err := runner.ReadIntGrid(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, list := range lists {
		if isSafePart1(list) {
			count += 1
		}
	}
	return count, nil
}

func problem2(lines []string) (int, error) {
	lists, err := runner.ReadIntGrid(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, list := range lists {
		if isSafePart1(list) {
			count += 1
		} else if isSafePart2(list) {
			count += 1
		}
	}
	return count, nil
}

func isSafePart1(list []int) bool {
	var incr bool
	for i := range list {
		if i == 0 {
			continue
		}
		if i == 1 {
			if list[i] > list[0] {
				incr = true
			}
		}
		if !incr && list[i-1] < list[i] {
			return false
		} else if incr && list[i-1] > list[i] {
			return false
		}

		diff := list[i] - list[i-1]
		if diff < 0 {
			diff = diff * -1
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafePart2(list []int) bool {
	for j := range list {
		list2 := make([]int, len(list)-1)
		copy(list2, list[:j])
		copy(list2[j:], list[j+1:])
		if isSafePart1(list2) {
			return true
		}
	}
	return false
}
