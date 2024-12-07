package main

import (
	"testing"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func TestProblem1(t *testing.T) {
	runner.RunTest(t, problem1, 3749)
}

func TestProblem2(t *testing.T) {
	runner.RunTest(t, problem2, 11387)
}

func TestCompute(t *testing.T) {
	for _, test := range []struct {
		key    int
		nums   []int
		concat bool
	}{
		{7290, []int{6, 8, 6, 15}, true},
		{192, []int{17, 8, 14}, true},
	} {
		if !compute(test.key, test.nums, test.concat) {
			t.Errorf("compute(%d, %v, %t) = false; want = true", test.key, test.nums, test.concat)
		}
	}
}
