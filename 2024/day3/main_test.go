package main

import (
	"testing"
)

func TestProblem2(t *testing.T) {
	test := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}
	if err := problem2(test); err != nil {
		t.Fatal(err)
	}
}
