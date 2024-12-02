package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const input = "inputs.txt"

func main() {
	list1, list2, err := createLists()
	if err != nil {
		log.Fatal(err)
	}
	problem1(list1, list2)
	problem2(list1, list2)
}

func problem1(list1, list2 []int) {
	var score int
	for i := range list1 {
		a := list1[i]
		b := list2[i]
		c := a - b
		if c < 0 {
			c = c * -1
		}
		score += c
	}
	fmt.Printf("Answer (part 1): %d\n", score)
}

func problem2(list1, list2 []int) {
	var score int
	for i := range list1 {
		for j := range list2 {
			if list1[i] == list2[j] {
				score += list1[i]
			}
		}
	}
	fmt.Printf("Answer (part 2): %d\n", score)
}

func createLists() (_ []int, _ []int, err error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	lists := make([][]int, 2)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		for i, n := range parts {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, nil, err
			}
			lists[i] = append(lists[i], num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	if len(lists[0]) != len(lists[1]) {
		return nil, nil, fmt.Errorf("lists[0] = %d; lists[1] = %d", len(lists[0]), lists[1])
	}

	sort.Ints(lists[0])
	sort.Ints(lists[1])
	return lists[0], lists[1], nil
}
