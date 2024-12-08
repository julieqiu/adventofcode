package main

import (
	"testing"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func TestProblem1(t *testing.T) {
	runner.RunTest(t, problem1, 14)
}

func TestProblem2(t *testing.T) {
	runner.RunTestWithFile(t, problem2, 9, "example5.txt")
}

func TestProblemA1(t *testing.T) {
	runner.RunTestWithFile(t, problem1, 2, "example2.txt")
}

func TestProblemA2(t *testing.T) {
	runner.RunTestWithFile(t, problem1, 4, "example3.txt")
}

func TestProblemA3(t *testing.T) {
	runner.RunTestWithFile(t, problem1, 4, "example4.txt")
}
