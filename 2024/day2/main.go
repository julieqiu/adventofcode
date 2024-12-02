package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const input = "input.txt"

func main() {
	lines, err := readLines()
	if err != nil {
		log.Fatal(err)
	}
	if err := problem1(lines); err != nil {
		log.Fatal(err)
	}
	if err := problem2(lines); err != nil {
		log.Fatal(err)
	}
}

func problem1(lists [][]int) error {
	var count int
	for _, list := range lists {
		if isSafePart1(list) {
			count += 1
		}
	}
	fmt.Printf("Answer (part 1): %d\n", count)
	return nil
}

func problem2(lists [][]int) error {
	var count int
	for _, list := range lists {
		if isSafePart1(list) {
			count += 1
		} else if isSafePart2(list) {
			count += 1
		}
	}
	fmt.Printf("Answer (part 2): %d\n", count)
	return nil
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

func truncateList(i int, list []int) []int {
	if i+1 >= len(list) {
		return list[:len(list)-1]
	}
	return append(list[:i], list[i+1:]...)
}

func readLines() (_ [][]int, err error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	var lists [][]int
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		var list []int
		for _, n := range parts {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			list = append(list, num)
		}
		lists = append(lists, list)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}
