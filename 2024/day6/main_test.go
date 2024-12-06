package main

import (
	"testing"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func TestProblem1(t *testing.T) {
	runner.RunTest(t, problem1, 41)
}

func TestProblem2(t *testing.T) {
	runner.RunTest(t, problem2, 6)
}

func TestInifiniteLoop(t *testing.T) {
	lines, err := runner.ReadLines("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	grid := runner.ReadRuneGrid(lines)
	curr := findGuard(grid)

	grid[9][81] = 'O'
	curr.x = 9
	curr.y = 20
	curr.dir = RIGHT
	result := countPositions(curr, grid)
	if result != -1 {
		t.Error(result)
	}
	printGrid(grid)
}

func Test923Loop(t *testing.T) {
	lines, err := runner.ReadLines("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	grid := runner.ReadRuneGrid(lines)
	curr := findGuard(grid)

	grid[9][23] = 'O'
	result := countPositions(curr, grid)
	if result != -1 {
		t.Error(result)
	}
	printGrid(grid)
}
