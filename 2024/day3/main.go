package main

import (
	"log"
	"strconv"
	"strings"

	runner "github.com/julieqiu/adventofcode/2024/internal/runner"
)

func main() {
	if err := runner.Run(problem1, problem2); err != nil {
		log.Fatal(err)
	}
}

func problem1(lines []string) (int, error) {
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
	return sum, nil
}

func problem2(lines []string) (int, error) {
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
	return sum, nil
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
