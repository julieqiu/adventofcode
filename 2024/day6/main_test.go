package main

import (
	"testing"

	"github.com/julieqiu/adventofcode/2024/internal/grid"
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

	g := grid.New(lines)
	curr := findGuard(g)

	g.Set(9, 81, 'O')
	curr.X = 9
	curr.Y = 20
	curr.dir = RIGHT
	result := countPositions(curr, g)
	if result != -1 {
		t.Error(result)
	}
	g.Print()
}

func Test923Loop(t *testing.T) {
	lines, err := runner.ReadLines("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	g := grid.New(lines)
	curr := findGuard(g)

	g.Set(9, 23, 'O')
	result := countPositions(curr, g)
	if result != -1 {
		t.Error(result)
	}
	g.Print()
}
