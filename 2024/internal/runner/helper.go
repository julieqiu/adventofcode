package runner

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Run(problem1, problem2 func([]string) (int, error)) error {
	input, err := readFile("input.txt")
	if err != nil {
		return err
	}
	lines, err := readLines(input)
	if err != nil {
		return err
	}
	ans, err := problem1(lines)
	if err != nil {
		return err
	}
	fmt.Printf("Answer (part 1): %v\n", ans)

	ans, err = problem2(lines)
	if err != nil {
		return err
	}
	fmt.Printf("Answer (part 2): %v\n", ans)
	return nil
}

func RunTest(t *testing.T, problem func([]string) (int, error), want int) {
	t.Helper()

	input, err := readFile("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	lines, err := readLines(input)
	if err != nil {
		t.Fatal(err)
	}
	got, err := problem(lines)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got = %d; want = %d", got, want)
	}
}

func readLines(input string) (_ []string, err error) {
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

func readFile(f string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", dir, f), nil
}
