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

func problem1(lines []string) error {
	var sum int
	for _, line := range lines {
		chunks := strings.Split(line, "mul(")
		for _, chunk := range chunks {
			part := strings.SplitN(chunk, ",", 2)
			a, err := strconv.Atoi(part[0])
			if err != nil {
				// invalid first number
				continue
			}

			parts2 := strings.Split(part[1], ")")
			b, err := strconv.Atoi(parts2[0])
			if err != nil {
				// invalid first number
				continue
			}
			sum += mul(a, b)
		}
	}
	fmt.Printf("Answer (part 1): %d\n", sum)
	return nil
}

func problem2(lines []string) error {
	var (
		sum  int
		skip bool
	)
	for _, line := range lines {
		i := 0
		for {
			if i > len(line)-1 {
				break
			}
			if strings.HasPrefix(line[i:], "mul(") {
				i += 4
				if skip {
					continue
				}
				result, j := maybeMultiply(line[i:])
				i += j
				sum += result
			} else if strings.HasPrefix(line[i:], "do") {
				const dont = "don't"
				if strings.HasPrefix(line[i:], dont) {
					i += len(dont)
					skip = true
					continue
				}
				i += 2
				skip = false
			} else {
				i += 1
			}
		}
	}
	fmt.Printf("Answer (part 2): %d\n", sum)
	return nil
}

func maybeMultiply(line string) (sum int, incr int) {
	parts := strings.SplitN(line, ",", 2)
	a, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0
	}
	parts2 := strings.Split(parts[1], ")")
	b, err := strconv.Atoi(parts2[0])
	if err != nil {
		return 0, 0
	}
	incr = len(parts[0]) + len(parts2[0])
	return mul(a, b), incr
}

func mul(a, b int) int {
	return a * b
}

func readLines() (_ []string, err error) {
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

	var list []string
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
