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

type point struct {
	x int
	y int
}

func problem1(lines []string) (int, error) {
	return solve(lines, false)
}

func problem2(lines []string) (int, error) {
	return solve(lines, true)
}

func solve(lines []string, includeHarmonics bool) (int, error) {
	antennas := map[rune][]point{}
	for y, row := range lines {
		for x, val := range row {
			if val != '.' {
				antennas[val] = append(antennas[val], point{x: x, y: y})
			}
		}
	}

	seen := map[point]bool{}
	for _, ants := range antennas {
		antinodes := findAntinodes(ants, len(lines[0]), len(lines), includeHarmonics)
		for _, at := range antinodes {
			seen[at] = true
		}
	}

	grid := runner.ReadRuneGrid(lines)
	for p := range seen {
		grid[p.y][p.x] = '#'
	}
	for i, row := range grid {
		fmt.Println(fmt.Sprintf("%.2d", i), string(row))
	}
	fmt.Println()
	return len(seen), nil
}

func findAntinodes(positions []point, limitx, limity int, includeHarmonics bool) []point {
	var output []point
	for i, p1 := range positions {
		for _, p2 := range positions[i:] {
			if p1 == p2 {
				continue
			}
			dx := p1.x - p2.x
			dy := p1.y - p2.y
			if dx < 0 {
				dx *= -1
			}
			if dy < 0 {
				dy *= -1
			}
			antinodes := possibleAntinodes(p1, p2, dx, dy, limitx, limity, includeHarmonics)
			output = append(output, antinodes...)
		}
	}
	return output
}

const (
	LEFT  = -1
	RIGHT = 1
	UP    = -1
	DOWN  = 1
)

func possibleAntinodes(a1, a2 point, dx, dy, limitx, limity int, includeHarmonics bool) []point {
	leftp := a1
	rightp := a2
	if a1.x > a2.x {
		leftp = a2
		rightp = a1
	}

	var positions []point
	if leftp.y < rightp.y {
		positions = append(positions, findAll(leftp, LEFT, UP, dx, dy, limitx, limity, includeHarmonics)...)
		positions = append(positions, findAll(rightp, RIGHT, DOWN, dx, dy, limitx, limity, includeHarmonics)...)
	} else {
		positions = append(positions, findAll(leftp, LEFT, DOWN, dx, dy, limitx, limity, includeHarmonics)...)
		positions = append(positions, findAll(rightp, RIGHT, UP, dx, dy, limitx, limity, includeHarmonics)...)
	}
	if includeHarmonics {
		positions = append(positions, leftp, rightp)
	}
	return positions
}

func findAll(p point, dirx, diry, dx, dy, limitx, limity int, includeHarmonics bool) (result []point) {
	for {
		dx2 := dx * dirx
		dy2 := dy * diry
		a1 := antinode(p, dx2, dy2)
		if outOfBounds(a1, limitx, limity) {
			return result
		}
		result = append(result, a1)
		if !includeHarmonics {
			return result
		}
		p = a1
	}
}

func outOfBounds(p point, limitx, limity int) bool {
	if p.x < 0 || p.y < 0 {
		return true
	}
	if p.x >= limitx || p.y >= limity {
		return true
	}
	return false
}

func antinode(p point, deltax, deltay int) point {
	return point{
		x: p.x + deltax,
		y: p.y + deltay,
	}
}
