package main

import (
	"fmt"
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
	g := grid.New(lines)
	start := findGuard(g)
	count := countPositions(start, g)
	return count, nil
}

func problem2(lines []string) (int, error) {
	g := grid.New(lines)
	start := findGuard(g)
	gridCopy := g.Copy()

	var count int
	g.Walk(func(r, c int) {
		newgrid := gridCopy.Copy()
		newgrid.Set(r, c, 'O')

		result := countPositions(start, newgrid)
		if result == -1 {
			count += 1
			fmt.Println(r, c)
		}
	})
	return count, nil
}

var (
	UP    = direction{Point: grid.Point{X: grid.UP, Y: 0}}
	DOWN  = direction{Point: grid.Point{X: grid.DOWN, Y: 0}}
	LEFT  = direction{Point: grid.Point{X: 0, Y: grid.LEFT}}
	RIGHT = direction{Point: grid.Point{X: 0, Y: grid.RIGHT}}
)

type position struct {
	grid.Point
	dir direction
}

type direction struct {
	grid.Point
}

func findGuard(g *grid.Grid) position {
	var curr position
	g.Walk(func(r, c int) {
		val := g.Get(r, c)
		switch val {
		case '^':
			curr = position{Point: grid.Point{X: r, Y: c}, dir: UP}
		case '>':
			curr = position{Point: grid.Point{X: r, Y: c}, dir: RIGHT}
		case 'v':
			curr = position{Point: grid.Point{X: r, Y: c}, dir: DOWN}
		case '<':
			curr = position{Point: grid.Point{X: r, Y: c}, dir: LEFT}
		}
	})
	return curr
}

func countPositions(curr position, g *grid.Grid) int {
	c := 0
	seen := map[position]bool{}
	for {
		c += 1
		seen[curr] = true

		row := curr.Y + curr.dir.Y
		col := curr.X + curr.dir.X

		if !g.InBounds(col, row) {
			break
		}

		if g.Equal(col, row, '#') || g.Equal(col, row, 'O') {
			curr.dir = turn(curr.dir)
			if !g.Equal(curr.X, curr.Y, '^') {
				g.Set(curr.X, curr.Y, '+')
			}
			continue
		}

		// Advance forward.
		curr.X = col
		curr.Y = row
		if seen[curr] {
			return -1
		}

		if !g.Equal(col, row, '^') {
			if curr.dir.X != 0 {
				if g.Equal(col, row, '-') {
					g.Set(col, row, '+')
				} else {
					g.Set(col, row, '|')
				}
			} else {
				if g.Equal(col, row, '|') {
					g.Set(col, row, '+')
				} else {
					g.Set(col, row, '-')
				}
			}
		}
	}

	visit := map[grid.Point]struct{}{}
	for s := range seen {
		visit[grid.Point{X: s.X, Y: s.Y}] = struct{}{}
	}
	return len(visit)
}

func turn(dir direction) direction {
	swap := dir.X * -1
	dir.X = dir.Y
	dir.Y = swap
	return dir
}
