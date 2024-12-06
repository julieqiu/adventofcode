package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/julieqiu/adventofcode/2024/internal/runner"
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
	runner.RunTest(t, problem1, 18)

	input := strings.Fields(testInput)
	ws := newWordSearch(input)
	if diff := cmp.Diff(want, ws.printUsed()); diff != "" {
		t.Log(diff)
	}
}

func TestProblem2(t *testing.T) {
	runner.RunTest(t, problem2, 9)
}
