package grid

import "fmt"

const (
	UP    = -1
	DOWN  = 1
	LEFT  = -1
	RIGHT = 1
)

type Grid struct {
	grid [][]rune
}

type Point struct {
	X int
	Y int
}

func New(lines []string) *Grid {
	g := &Grid{
		grid: make([][]rune, len(lines)),
	}
	for r, line := range lines {
		g.grid[r] = make([]rune, len(lines[0]))
		for c, cell := range line {
			g.grid[r][c] = cell
		}
	}
	return g
}

func (g *Grid) Print() {
	for i, row := range g.grid {
		fmt.Printf("%.2d %s\n", i, string(row))
	}
	fmt.Println()
}

func (g *Grid) PrintWithCol() {
	fmt.Print("   ")
	for j := range g.grid[0] {
		fmt.Printf("%.2d ", j)
	}
	fmt.Println()
	for i, row := range g.grid {
		fmt.Printf("%.2d", i)
		for _, val := range row {
			fmt.Printf("%3c", val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Grid) Copy() *Grid {
	newgrid := make([][]rune, len(g.grid))
	for r := range newgrid {
		newgrid[r] = make([]rune, len(g.grid[0]))
		for c := range newgrid {
			newgrid[r][c] = g.grid[r][c]
		}
	}
	return &Grid{
		grid: newgrid,
	}
}
func (g *Grid) Get(row, col int) rune {
	return g.grid[row][col]
}

func (g *Grid) Set(row, col int, v rune) {
	g.grid[row][col] = v
}

func (g *Grid) Equal(row, col int, v rune) bool {
	return g.grid[row][col] == v
}

func (g *Grid) Walk(fn func(r, c int)) {
	for r, row := range g.grid {
		for c := range row {
			fn(r, c)
		}
	}
}

func (g *Grid) InBounds(r, c int) bool {
	if c < 0 || c >= len(g.grid) {
		return false
	}
	if r < 0 || r >= len(g.grid[0]) {
		return false
	}
	return true
}
