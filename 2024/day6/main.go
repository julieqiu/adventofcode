package main

import (
	"fmt"
	"log"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func main() {
	if err := runner.Run(problem1, problem2); err != nil {
		log.Fatal(err)
	}
}

func problem1(lines []string) (int, error) {
	grid := runner.ReadRuneGrid(lines)
	start := findGuard(grid)
	count := countPositions(start, grid)
	return count, nil
}

func problem2(lines []string) (int, error) {
	grid := runner.ReadRuneGrid(lines)
	start := findGuard(grid)
	gridCopy := newGrid(grid)

	var count int
	for x := range grid {
		for y := range grid[0] {
			newgrid := newGrid(gridCopy)
			newgrid[x][y] = 'O'

			result := countPositions(start, newgrid)
			if result == -1 {
				count += 1
				fmt.Println(x, y)
			}
		}
	}
	return count, nil
}

var (
	UP    = direction{x: -1, y: 0}
	DOWN  = direction{x: 1, y: 0}
	LEFT  = direction{x: 0, y: -1}
	RIGHT = direction{x: 0, y: 1}
)

type position struct {
	x   int
	y   int
	dir direction
}

type direction struct {
	x int
	y int
}

func findGuard(grid [][]rune) position {
	var curr position
	for r, row := range grid {
		for c, val := range row {
			switch val {
			case '^':
				curr = position{x: r, y: c, dir: UP}
			case '>':
				curr = position{x: r, y: c, dir: RIGHT}
			case 'v':
				curr = position{x: r, y: c, dir: DOWN}
			case '<':
				curr = position{x: r, y: c, dir: LEFT}
			}
		}
	}
	return curr
}

func countPositions(curr position, grid [][]rune) int {
	c := 0
	seen := map[position]bool{}
	for {
		c += 1
		seen[curr] = true

		x := curr.x + curr.dir.x
		y := curr.y + curr.dir.y

		if x < 0 || x >= len(grid) {
			break
		}
		if y < 0 || y >= len(grid[0]) {
			break
		}

		if grid[x][y] == '#' || grid[x][y] == 'O' {
			curr.dir = turn(curr.dir)
			if grid[curr.x][curr.y] != '^' {
				grid[curr.x][curr.y] = '+'
			}
			continue
		}

		// Advance forward.
		curr.x = x
		curr.y = y
		if seen[curr] {
			return -1
		}

		if grid[x][y] != '^' {
			if curr.dir.x != 0 {
				if grid[x][y] == '-' {
					grid[x][y] = '+'
				} else {
					grid[x][y] = '|'
				}
			} else {
				if grid[x][y] == '|' {
					grid[x][y] = '+'
				} else {
					grid[x][y] = '-'
				}
			}
		}
	}

	visit := map[position]struct{}{}
	for s := range seen {
		visit[position{x: s.x, y: s.y}] = struct{}{}
	}
	return len(visit)
}

func turn(dir direction) direction {
	swap := dir.x * -1
	dir.x = dir.y
	dir.y = swap
	return dir
}
