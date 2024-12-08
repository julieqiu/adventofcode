package main

import (
	"log"

	"github.com/julieqiu/adventofcode/2024/internal/grid"
	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func main() {
	if err := runner.Run(problem1, problem2); err != nil {
		log.Fatal(err)
	}
}

func problem1(lines []string) (int, error) {
	ws := newWordSearch(lines)
	count := ws.solveXMAS()
	return count, nil
}

func problem2(lines []string) (int, error) {
	ws := newWordSearch(lines)
	count := ws.solveMAS()
	return count, nil
}

type wordsearch struct {
	lines []string
	used  [][]bool
	count int
}

func newWordSearch(lines []string) *wordsearch {
	ws := &wordsearch{lines: lines}
	ws.used = make([][]bool, len(lines))
	for i := range ws.used {
		ws.used[i] = make([]bool, len(lines[0]))
	}
	return ws
}

func (ws *wordsearch) printUsed() string {
	var output string
	for i, row := range ws.used {
		var toPrint string
		for j, used := range row {
			if used {
				toPrint += string(ws.lines[i][j])
			} else {
				toPrint += "."
			}
		}
		output += toPrint + "\n"
	}
	return "\n" + output
}

const (
	XMAS     = "XMAS"
	MAS      = "MAS"
	STRAIGHT = 0
)

func (ws *wordsearch) solveMAS() int {
	for row, line := range ws.lines {
		for col := range line {
			if ws.search(MAS, row+grid.LEFT, col+grid.UP, grid.RIGHT, grid.DOWN)+
				ws.search(MAS, row+grid.LEFT, col+grid.DOWN, grid.RIGHT, grid.UP)+
				ws.search(MAS, row+grid.RIGHT, col+grid.UP, grid.LEFT, grid.DOWN)+
				ws.search(MAS, row+grid.RIGHT, col+grid.DOWN, grid.LEFT, grid.UP) >= 2 {
				ws.count += 1
			}
		}
	}
	return ws.count
}

func (ws *wordsearch) solveXMAS() int {
	for row, line := range ws.lines {
		for col, c := range line {
			if string(c) != "X" {
				continue
			}
			ws.count += ws.search(XMAS, row, col, grid.RIGHT, STRAIGHT)
			ws.count += ws.search(XMAS, row, col, grid.LEFT, STRAIGHT)
			ws.count += ws.search(XMAS, row, col, STRAIGHT, grid.UP)
			ws.count += ws.search(XMAS, row, col, STRAIGHT, grid.DOWN)
			ws.count += ws.search(XMAS, row, col, grid.RIGHT, grid.UP)
			ws.count += ws.search(XMAS, row, col, grid.RIGHT, grid.DOWN)
			ws.count += ws.search(XMAS, row, col, grid.LEFT, grid.UP)
			ws.count += ws.search(XMAS, row, col, grid.LEFT, grid.DOWN)
		}
	}
	return ws.count
}

func (ws *wordsearch) search(input string, row, col, h, v int) int {
	for k := range input {
		row2 := row + h*k
		col2 := col + v*k

		if row2 < 0 || row2 >= len(ws.lines) {
			return 0
		}
		if col2 < 0 || col2 >= len(ws.lines[0]) {
			return 0
		}
		if ws.lines[row2][col2] != input[k] {
			return 0
		}
	}
	for k := range input {
		ws.used[row+h*k][col+v*k] = true
	}
	return 1
}
