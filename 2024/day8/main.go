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
	return solve(lines, false)
}

func problem2(lines []string) (int, error) {
	return solve(lines, true)
}

func solve(lines []string, includeHarmonics bool) (int, error) {
	antennas := map[rune][]grid.Point{}
	g := grid.New(lines)
	g.Walk(func(r, c int) {
		val := g.Get(r, c)
		if val != '.' {
			antennas[val] = append(antennas[val], grid.Point{X: c, Y: r})
		}
	})

	seen := map[grid.Point]bool{}
	for _, ants := range antennas {
		antinodes := findAntinodesForSymbol(ants, len(lines[0]), len(lines), includeHarmonics)
		for _, at := range antinodes {
			seen[at] = true
		}
	}

	for p := range seen {
		g.Set(p.Y, p.X, '#')
	}
	g.Print()
	return len(seen), nil
}

func findAntinodesForSymbol(antennas []grid.Point, limitx, limity int, includeHarmonics bool) []grid.Point {
	var output []grid.Point
	for i, p1 := range antennas {
		for _, p2 := range antennas[i:] {
			if p1 == p2 {
				continue
			}
			dx := p1.X - p2.X
			dy := p1.Y - p2.Y
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

func possibleAntinodes(p1, p2 grid.Point, dx, dy, limitx, limity int, includeHarmonics bool) []grid.Point {
	leftp := p1
	rightp := p2
	if p1.X > p2.X {
		leftp = p2
		rightp = p1
	}

	var antinodes []grid.Point
	if includeHarmonics {
		antinodes = append(antinodes, leftp, rightp)
	}
	if leftp.Y < rightp.Y {
		antinodes = append(antinodes, findAll(leftp, grid.LEFT, grid.UP, dx, dy, limitx, limity, includeHarmonics)...)
		antinodes = append(antinodes, findAll(rightp, grid.RIGHT, grid.DOWN, dx, dy, limitx, limity, includeHarmonics)...)
	} else {
		antinodes = append(antinodes, findAll(leftp, grid.LEFT, grid.DOWN, dx, dy, limitx, limity, includeHarmonics)...)
		antinodes = append(antinodes, findAll(rightp, grid.RIGHT, grid.UP, dx, dy, limitx, limity, includeHarmonics)...)
	}
	return antinodes
}

func findAll(p grid.Point, dirx, diry, dx, dy, limitx, limity int, includeHarmonics bool) (result []grid.Point) {
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

func outOfBounds(p grid.Point, limitx, limity int) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}
	if p.X >= limitx || p.Y >= limity {
		return true
	}
	return false
}

func antinode(p grid.Point, deltax, deltay int) grid.Point {
	return grid.Point{
		X: p.X + deltax,
		Y: p.Y + deltay,
	}
}
