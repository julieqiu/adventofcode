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
	g, err := runner.ReadIntGrid(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, start := range findTrailheads(g) {
		c := findSummit(start, g)
		count += c
	}
	return count, nil
}

func problem2(lines []string) (int, error) {
	g, err := runner.ReadIntGrid(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, start := range findTrailheads(g) {
		c := findDistinctTrails(start, g)
		count += c
	}
	return count, nil
}

func findTrailheads(g [][]int) (output []grid.Point) {
	for y, row := range g {
		for x, val := range row {
			if val == 0 {
				output = append(output, grid.Point{
					X: x,
					Y: y,
				})
			}
		}
	}
	return output
}

func findDistinctTrails(start grid.Point, g [][]int) int {
	var row, col, r1, r2, c1, c2, curr, count int
	var queue = []grid.Point{start}
	for {
		if len(queue) == 0 {
			break
		}
		row = queue[0].Y
		col = queue[0].X
		queue = queue[1:]

		curr = g[row][col]
		if curr+1 == 10 {
			count += 1
			continue
		}
		r1 = row + grid.UP
		r2 = row + grid.DOWN
		c1 = col + grid.LEFT
		c2 = col + grid.RIGHT

		for _, p := range []grid.Point{
			{Y: row, X: c1},
			{Y: row, X: c2},
			{Y: r1, X: col},
			{Y: r2, X: col},
		} {
			if outOfBounds(p.Y, p.X, g) {
				continue
			}
			if g[p.Y][p.X] == curr+1 {
				queue = append(queue, grid.Point{Y: p.Y, X: p.X})
			}
		}
	}
	return count

}

func findSummit(start grid.Point, g [][]int) int {
	var row, col, r1, r2, c1, c2, curr, count int
	var queue = []grid.Point{start}
	for {
		if len(queue) == 0 {
			break
		}
		row = queue[0].Y
		col = queue[0].X
		queue = queue[1:]

		curr = g[row][col]
		if curr+1 == 10 {
			count += 1
			continue
		}
		r1 = row + grid.UP
		r2 = row + grid.DOWN
		c1 = col + grid.LEFT
		c2 = col + grid.RIGHT

		for _, p := range []grid.Point{
			{Y: row, X: c1},
			{Y: row, X: c2},
			{Y: r1, X: col},
			{Y: r2, X: col},
		} {
			if outOfBounds(p.Y, p.X, g) {
				continue
			}
			if g[p.Y][p.X] == curr+1 {
				var seen bool
				for _, q := range queue {
					if q.Y == p.Y && q.X == p.X {
						seen = true
					}
				}
				if !seen {
					queue = append(queue, grid.Point{Y: p.Y, X: p.X})
				}
			}
		}
	}
	return count
}

func outOfBounds(row, col int, g [][]int) bool {
	if row < 0 || col < 0 {
		return true
	}
	if row > len(g)-1 || col > len(g[row])-1 {
		return true
	}
	return false
}
