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
	return parseBlinks(lines[0], 25)
}

func problem2(lines []string) (int, error) {
	return parseBlinksPart2(lines[0])
}

func parseBlinksPart2(line string) (int, error) {
	stones, err := parseStones(line)
	if err != nil {
		return 0, err
	}
	var count int
	for _, stone := range stones {
		c, err := countStones(stone, stone, 0)
		if err != nil {
			return 0, err
		}
		count += c
	}
	return count, nil
}

func countStones(stone, original, blinks int) (int, error) {
	if blinks == 75 {
		return 1, nil
	}
	output, err := convertStone(stone)
	if err != nil {
		return 0, err
	}
	var count int
	for _, n := range output {
		fmt.Printf("countStones(%d, %d, %d)\n", n, original, blinks+1)
		c, err := countStones(n, original, blinks+1)
		if err != nil {
			return 0, err
		}
		count += c
	}
	return count, nil
}

func parseBlinks(line string, blinks int) (int, error) {
	stones, err := parseStones(line)
	if err != nil {
		return 0, err
	}
	for i := 0; i < blinks; i++ {
		var newstones []int
		for _, stone := range stones {
			s, err := convertStone(stone)
			if err != nil {
				return 0, err
			}
			newstones = append(newstones, s...)
		}
		stones = newstones
	}
	return len(stones), nil
}

func parseStones(line string) ([]int, error) {
	var stones []int
	parts := strings.Fields(line)
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		stones = append(stones, n)
	}
	return stones, nil
}

func convertStone(stone int) (output []int, err error) {
	// If the stone is engraved with the number 0, it is replaced by a stone
	// engraved with the number 1.
	if stone == 0 {
		return []int{1}, nil
	}

	// If the stone is engraved with a number that has an even number of digits, it
	// is replaced by two stones. The left half of the digits are engraved on the
	// new left stone, and the right half of the digits are engraved on the new
	// right stone. (The new numbers don't keep extra leading zeroes: 1000 would
	// become stones 10 and 0.)
	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		s1 := s[0 : len(s)/2]
		n1, err := strconv.Atoi(s1)
		if err != nil {
			return nil, err
		}
		output = append(output, n1)

		s2 := s[len(s)/2:]
		n2, err := strconv.Atoi(s2)
		if err != nil {
			return nil, err
		}
		output = append(output, n2)
		return output, nil
	}

	// If none of the other rules apply, the stone is replaced by a new stone;
	// the old stone's number multiplied by 2024 is engraved on the new stone.
	return []int{stone * 2024}, nil
}
