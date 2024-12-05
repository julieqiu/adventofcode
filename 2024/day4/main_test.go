package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const testInput = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

const want = `
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
`

func TestProblem1(t *testing.T) {
	input := strings.Fields(testInput)
	ws := newWordSearch(input)
	ans := ws.solveXMAS()
	if diff := cmp.Diff(want, ws.printUsed()); diff != "" {
		fmt.Println(diff)
		fmt.Println(ans)
	}
}

func TestProblem2(t *testing.T) {
	input := strings.Fields(testInput)
	ws := newWordSearch(input)
	ans := ws.solveMAS()
	if ans != 9 {
		t.Error(ans)
	}
}
