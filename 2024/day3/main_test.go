package main

import (
	"testing"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func TestProblem1(t *testing.T) {
	runner.RunTest(t, problem1, 161)
}

func TestProblem2(t *testing.T) {
	runner.RunTest(t, problem2, 48)
}
