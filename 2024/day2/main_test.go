package main

import (
	"fmt"
	"testing"
)

func TestIsSafe(t *testing.T) {
	for _, test := range []struct {
		input []int
		want1 bool
		want2 bool
	}{
		{
			[]int{7, 6, 4, 2, 1},
			true,
			true,
		},
		{
			[]int{1, 2, 7, 8, 9},
			false,
			false,
		},
		{
			[]int{9, 7, 6, 2, 1},
			false,
			false,
		},
		{
			[]int{1, 3, 2, 4, 5},
			false,
			true,
		},
		{
			[]int{8, 6, 4, 4, 1},
			false,
			true,
		},
		{
			[]int{1, 3, 6, 7, 9},
			true,
			true,
		},
	} {
		t.Run(fmt.Sprintf("part 1 %v", test.input), func(t *testing.T) {
			if got := isSafePart1(test.input); got != test.want1 {
				t.Errorf("want = %t; got = %t", got, test.want1)
			}
		})
		t.Run(fmt.Sprintf("part 2 %v", test.input), func(t *testing.T) {
			if got := isSafePart2(test.input); got != test.want2 {
				t.Errorf("want = %t; got = %t", got, test.want2)
			}
		})
	}
}
