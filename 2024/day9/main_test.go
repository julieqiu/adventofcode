package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func TestProblem1(t *testing.T) {
	runner.RunTest(t, problem1, 1928)
}

func TestProblem2(t *testing.T) {
	runner.RunTest(t, problem2, 2858)
}

func TestFindSpace(t *testing.T) {
	test := "00...111...2...333.44.5555.6666.777.888899"
	want := [][]int{
		{2, 5},
		{8, 11},
		{12, 15},
		{18, 19},
		{21, 22},
		{26, 27},
		{31, 32},
		{35, 36},
	}
	var array []string
	for _, s := range test {
		array = append(array, string(s))
	}
	var start, end int
	for start != -1 {
		start, end = findSpace(end, array)
		if start == -1 {
			break
		}
		w := want[0]
		if start != w[0] && end != w[1] {
			t.Fatalf("got = %d %d; want = %d %d", start, end, w[0], w[1])
		}
		want = want[1:]
	}
}

func TestFindBlock(t *testing.T) {
	test := "00...111...2...333.44.5555.6666.777.888899"
	want := [][]int{
		{9, 40, 42},
		{8, 36, 40},
		{7, 32, 35},
		{6, 27, 31},
		{5, 24, 26},
		{4, 19, 21},
		{3, 15, 18},
		{2, 11, 12},
		{1, 5, 8},
		{0, 0, 2},
	}
	var array []string
	for _, s := range test {
		array = append(array, string(s))
	}

	start := len(array)
	end := len(array)
	val := 0
	for start != -1 {
		start, end, val = findBlock(val-1, start, array)
		if start == -1 {
			break
		}
		w := want[0]
		if val != w[0] && start != w[1] && end != w[2] {
			t.Fatalf("got = %d %d %d; want = %d %d %d", start, end, val, w[1], w[2], w[0])
		}
		want = want[1:]
	}
}

func TestRemoveFreeSpacePart2(t *testing.T) {
	for _, test := range []struct {
		input string
		want  string
	}{
		{"2333133121414131402", "00992111777.44.333....5555.6666.....8888.."},
	} {
		t.Run(test.input, func(t *testing.T) {
			g := removeFreeSpacePart2(test.input)
			got := strings.Join(g, "")
			if got != test.want {
				t.Errorf("got = %q; want = %q", got, test.want)
			}
		})
	}
}

func TestRemoveFreeSpace(t *testing.T) {
	for _, test := range []struct {
		input string
		want  string
	}{
		{"12345", "022111222......"},
		{"2333133121414131402", "0099811188827773336446555566.............."},
	} {
		t.Run(test.input, func(t *testing.T) {
			g := removeFreeSpace(test.input)
			got := strings.Join(g, "")
			if got != test.want {
				t.Errorf("got = %q; want = %q", got, test.want)
			}
		})
	}
}

func TestDiskmapToBlocks(t *testing.T) {
	for _, test := range []struct {
		input string
		want  string
	}{
		{"12345", "0..111....22222"},
		{"2333133121414131402", "00...111...2...333.44.5555.6666.777.888899"},
	} {
		t.Run(test.input, func(t *testing.T) {
			g := diskmapToBlocks(test.input)
			got := strings.Join(g, "")
			if got != test.want {
				t.Errorf("got = %q; want = %q", got, test.want)
			}
		})
	}
}

func toString(blocks []rune) (output []string) {
	for _, b := range blocks {
		fmt.Println(string(b), b)
		output = append(output, string(b))
	}
	return output
}

/*
func TestRune(t *testing.T) {
	// rune represents a character in Unicode.
	// The character 3 is 51 in Unicode.
	r := '3' // 51

	// i represents an integer.
	// The integer 3 is a non-printing character in Unicode.
	i := 3

	// s represents a character in bytes.
	// "3" and byte(51) are the same.
	s := "3"
	b2 := byte(51) // 3

	// byte(3) represents a non-printing character in Unicode.
	b := byte(3)

	t.Log(byte(51))
}
*/
